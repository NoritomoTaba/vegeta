package vegeta

import (
	"context"
	"crypto/subtle"
	"fmt"
	"net/http"

	"github.com/Code-Hex/vegeta/internal/utils"
	"github.com/Code-Hex/vegeta/model"
	"github.com/Code-Hex/vegeta/protos"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type API struct {
	DB *gorm.DB
}

const (
	week uint = iota + 1
	month
	all
)

/* grpc */
func (v *Vegeta) NewAPI() *API {
	return &API{DB: v.DB}
}

func (a *API) AddData(ctx context.Context, r *protos.RequestFromDevice) (*protos.ResultResponse, error) {
	token := r.GetToken()
	user, err := model.TokenAuth(a.DB, token)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	tag, err := user.FindByTagName(a.DB, r.GetTagName())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	data := model.Data{
		RemoteAddr: r.GetRemoteAddr(),
		Payload:    r.GetPayload(),
		Hostname:   r.GetHostname(),
	}
	if err := tag.AddData(a.DB, data); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &protos.ResultResponse{}, nil
}

func (a *API) AddTag(ctx context.Context, r *protos.AddTagFromDevice) (*protos.ResultResponse, error) {
	token := r.GetToken()
	user, err := model.TokenAuth(a.DB, token)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	tag := r.GetTagName()
	if _, err := user.FindByTagName(a.DB, tag); err == nil {
		return nil, status.Error(
			codes.NotFound,
			fmt.Sprintf("Tag: %s is already exists", tag),
		)
	}
	if err := user.AddTag(a.DB, tag); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &protos.ResultResponse{}, nil
}

/* Public JSON API */
type resultGetTagList struct {
	Tags []string `json:"tags"`
}

func GetTagList() echo.HandlerFunc {
	return call(func(c *Context) error {
		user, ok := c.Get("user").(*model.User)
		if !ok {
			return errors.New("Failed to get user info via context")
		}
		tags := make([]string, len(user.Tags), len(user.Tags))
		for i, v := range user.Tags {
			tags[i] = v.Name
		}
		return c.JSON(http.StatusOK, &resultGetTagList{
			Tags: tags,
		})
	})
}

type getDataList struct {
	Tag   string `json:"tag" validate:"required"`
	Span  string `json:"span" validate:"required"`
	Limit uint   `json:"limit" validate:"required"`
	Page  uint   `json:"page"`
}

type resultGetDataList struct {
	Data []model.Data `json:"data"`
}

func GetDataList() echo.HandlerFunc {
	return call(func(c *Context) error {
		param := new(getDataList)
		if err := c.BindValidate(param); err != nil {
			return err
		}
		user, ok := c.Get("user").(*model.User)
		if !ok {
			return errors.New("Failed to get user info via context")
		}
		tag, err := user.FindByTagName(c.DB, param.Tag)
		if err != nil {
			return errors.Wrap(err, "Failed to get tag")
		}
		page := param.Page
		span := param.Span
		limit := param.Limit
		data, err := model.FindDataByTagID(c.DB, tag.ID, page, limit, span)
		if err != nil {
			return errors.Wrap(err, "Failed to find data")
		}
		return c.JSON(http.StatusOK, &resultGetDataList{
			Data: data,
		})
	})
}

/* JSON API for settings */
type apiVegetaClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type resultJSON struct {
	IsSuccess bool   `json:"is_success"`
	Reason    string `json:"reason"`
}

func RegenerateToken() echo.HandlerFunc {
	return call(func(c *Context) error {
		token, ok := c.Get("auth_api").(*jwt.Token)
		if !ok {
			c.Zap.Info("Failed to check user has a permission")
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: "APIトークンにユーザーの情報がありませんでした",
			})
		}
		claim := token.Claims.(*apiVegetaClaims)
		user, err := model.FindUserByName(c.DB, claim.Name)
		if err != nil {
			c.Zap.Info("Failed to get user at /regenerate")
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: "トークンの更新に失敗しました",
			})
		}
		if _, err := user.ReGenerateUserToken(c.DB); err != nil {
			c.Zap.Info("Failed to regenerate token at /regenerate", zap.Error(err))
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: "トークンの更新に失敗しました",
			})
		}
		return c.JSON(http.StatusOK, &resultJSON{
			IsSuccess: true,
		})
	})
}

type reregisterPassword struct {
	Password       string `json:"password" validate:"required"`
	VerifyPassword string `json:"verify_password" validate:"required"`
}

func ReRegisterPassword() echo.HandlerFunc {
	return call(func(c *Context) error {
		param := new(reregisterPassword)
		if err := c.BindValidate(param); err != nil {
			return err
		}
		password := param.Password
		verifyPassword := param.VerifyPassword
		if subtle.ConstantTimeCompare([]byte(password), []byte(verifyPassword)) != 1 {
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: "入力したパスワードと確認用のパスワードが一致しませんでした。",
			})
		}
		token, ok := c.Get("auth_api").(*jwt.Token)
		if !ok {
			c.Zap.Info("Failed to check user has a permission")
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: "ユーザーの情報がありませんでした",
			})
		}
		claim := token.Claims.(*apiVegetaClaims)
		user, err := model.FindUserByName(c.DB, claim.Name)
		if err != nil {
			c.Zap.Info("Failed to get user at /reregister_password", zap.Error(err))
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: err.Error(),
			})
		}
		if _, err := user.UpdatePassword(c.DB, password); err != nil {
			c.Zap.Info("Failed to get user at /reregister_password", zap.Error(err))
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, &resultJSON{
			IsSuccess: true,
		})
	})
}

/* JSON API for mypage */
type addTag struct {
	Name string `json:"tag_name" validate:"required"`
}

func AddTag() echo.HandlerFunc {
	return call(func(c *Context) error {
		param := new(addTag)
		if err := c.BindValidate(param); err != nil {
			return err
		}
		token, ok := c.Get("auth_api").(*jwt.Token)
		if !ok {
			c.Zap.Info("Failed to check user has a permission")
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: "APIトークンにユーザーの情報がありませんでした",
			})
		}
		claim := token.Claims.(*apiVegetaClaims)
		user, err := model.FindUserByName(c.DB, claim.Name)
		if err != nil {
			c.Zap.Info("Failed to get user at /regenerate")
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: "トークンの更新に失敗しました",
			})
		}

		if err := user.AddTag(c.DB, param.Name); err != nil {
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, &resultJSON{
			IsSuccess: true,
		})
	})
}

type getTagsData struct {
	TagID uint   `json:"tag_id" validate:"required"`
	Span  string `json:"span" validate:"required"`
	Limit uint   `json:"limit" validate:"required"`
	Page  uint   `json:"page"`
}

type resultGetTagsJSON struct {
	IsSuccess bool         `json:"is_success"`
	Data      []model.Data `json:"data"`
}

func JSONTagsData() echo.HandlerFunc {
	return call(func(c *Context) error {
		param := new(getTagsData)
		if err := c.BindValidate(param); err != nil {
			return err
		}
		tagID := param.TagID
		page := param.Page
		span := param.Span
		limit := param.Limit
		data, err := model.FindDataByTagID(c.DB, tagID, page, limit, span)
		if err != nil {
			c.Zap.Info("Failed to get tag",
				zap.Error(err),
				zap.Uint("tag_id", param.TagID),
			)
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: "データを取得するときにエラーが発生しました",
			})
		}

		return c.JSON(http.StatusOK, &resultGetTagsJSON{
			IsSuccess: true,
			Data:      data,
		})
	})
}

/* JSON API for admin */
type createUser struct {
	Name           string `json:"name" validate:"required"`
	Password       string `json:"password" validate:"required"`
	VerifyPassword string `json:"verify_password" validate:"required"`
	IsAdmin        bool   `json:"is_admin"`
}

func JSONCreateUser() echo.HandlerFunc {
	return call(func(c *Context) error {
		param := new(createUser)
		if err := c.BindValidate(param); err != nil {
			return err
		}

		password := param.Password
		verifyPassword := param.VerifyPassword
		if subtle.ConstantTimeCompare([]byte(password), []byte(verifyPassword)) != 1 {
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: "入力したパスワードと確認用のパスワードが一致しませんでした。",
			})
		}
		username := param.Name
		isAdmin := param.IsAdmin
		if _, err := model.CreateUser(c.DB, username, password, isAdmin); err != nil {
			c.Zap.Error("Failed to create user", zap.Error(err))
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: "ユーザー作成時にエラーが発生しました。",
			})
		}
		return c.JSON(http.StatusOK, &resultJSON{
			IsSuccess: true,
		})
	})
}

type editUser struct {
	ID              string `json:"id" validate:"required"`
	IsAdmin         bool   `json:"is_admin"`
	IsResetPassword bool   `json:"is_reset_password"`
}

func JSONEditUser() echo.HandlerFunc {
	return call(func(c *Context) error {
		editUser := new(editUser)
		if err := c.BindValidate(editUser); err != nil {
			return err
		}

		userID := editUser.ID
		isAdmin := editUser.IsAdmin
		isResetPassword := editUser.IsResetPassword

		var str string
		if isResetPassword {
			str = utils.RandomString()
		}

		if _, err := model.EditUser(c.DB, userID, isAdmin, str); err != nil {
			c.Zap.Error("Failed to edit user", zap.Error(err))
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: "ユーザー編集時にエラーが発生しました。",
			})
		}
		if isResetPassword {
			return c.JSON(http.StatusOK, &resultJSON{
				IsSuccess: true,
				Reason:    str,
			})
		}
		return c.JSON(http.StatusOK, &resultJSON{
			IsSuccess: true,
		})
	})
}

type deleteUser struct {
	ID string `json:"id" validate:"required"`
}

func JSONDeleteUser() echo.HandlerFunc {
	return call(func(c *Context) error {
		deleteUser := new(deleteUser)
		if err := c.BindValidate(deleteUser); err != nil {
			return err
		}

		userID := deleteUser.ID
		if _, err := model.DeleteUser(c.DB, userID); err != nil {
			c.Zap.Error("Failed to delete user", zap.Error(err))
			return c.JSON(http.StatusOK, &resultJSON{
				Reason: "ユーザー削除時にエラーが発生しました。",
			})
		}
		return c.JSON(http.StatusOK, &resultJSON{
			IsSuccess: true,
		})
	})
}
