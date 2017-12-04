// Code generated by hero.
// source: /Users/codehex/Desktop/go/src/github.com/Code-Hex/vegeta/template/mypage.html
// DO NOT EDIT!
package html

import (
	"io"

	"github.com/shiyanhui/hero"
)

func MyPage(args MyPageArgs, w io.Writer) {
	_buffer := hero.GetBuffer()
	defer hero.PutBuffer(_buffer)
	_buffer.WriteString(`<!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta name="description" content="IoTを用いた栽培中の植物のデータを管理するプロジェクトです。">
  <link href="/assets/css/main.css" rel="stylesheet">
  <link href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css" rel="stylesheet" integrity="sha384-wvfXpqpZZVQGK6TAh5PVlGOfQNHSoD2xbE+QkPxCAFlNEevoEH3Sl0sibVcOQVnN" crossorigin="anonymous">
  <link rel="stylesheet" href="/assets/css/bootstrap.css">
  <script src="/assets/js/jquery.min.js"></script>
  <script src="/assets/js/tether.min.js"></script>
  <script src="/assets/js/bootstrap.min.js"></script>
  `)
	_buffer.WriteString(`
  <link rel="stylesheet" type="text/css" href="/assets/css/c3.min.css">
`)

	_buffer.WriteString(`
  <title>`)
	_buffer.WriteString(`mypage`)

	_buffer.WriteString(`</title>
</head>
<body class="d-flex flex-column" style="min-height: 100vh">
  <nav class="navbar navbar-toggleable-md navbar-expand-lg navbar-light static-top v-navbar">
    <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
      <i class="fa fa-bars"></i>
    </button>
    <a class="navbar-brand" href="/">Vegeta</a>
    <div id="navbarResponsive" class="collapse navbar-collapse">
      <ul class="navbar-nav mr-auto">
        <li class="nav-item"><a class="nav-link" href="/contact">問い合わせ</a></li>
      </ul>
      <ul class="navbar-nav">
        `)
	if args.IsAuthed() {
		_buffer.WriteString(`
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle dropdown-toggle-split" href="" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"><i class="fa fa-user" aria-hidden="true"></i> ユーザー</a>
            <div class="dropdown-menu">
              <a class="dropdown-item" href="/mypage"><i class="fa fa-pagelines" aria-hidden="true"></i> 観察</a>
              <div class="dropdown-divider"></div>
              <a class="dropdown-item" href="/mypage/settings"><i class="fa fa-cog" aria-hidden="true"></i> 設定</a>
              `)
		if args.IsAdmin() {
			_buffer.WriteString(`
                <a class="dropdown-item" href="/mypage/admin"><i class="fa fa-lock" aria-hidden="true"></i> ユーザー管理パネル</a>
              `)
		}
		_buffer.WriteString(`
            </div>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="/mypage/logout"><i class="fa fa-sign-out" aria-hidden="true"></i> ログアウト</a>
          </li>
        `)
	} else {
		_buffer.WriteString(`
          <li class="nav-item">
            <a class="nav-link" href="/login"><i class="fa fa-sign-in" aria-hidden="true"></i> ログイン</a>
          </li>
        `)
	}
	_buffer.WriteString(`
      </ul>
    </div>
  </nav>
  <main class="mb-auto">
    `)

	mypageArgs := args
	user := mypageArgs.User()

	_buffer.WriteString(`
<input type="hidden" id="api-token" value="`)
	hero.EscapeHTML(mypageArgs.Token(), _buffer)
	_buffer.WriteString(`">
<div class="content">
  <div class="container">
    `)
	if len(user.Tags) > 0 {
		_buffer.WriteString(`
      <div class="row float-right">
        <div class="col">
          <select id="action" class="form-control tag-select">
            <option value="">タグ一覧</option>
            `)
		for _, tag := range user.Tags {
			_buffer.WriteString(`
              <option value="`)
			hero.FormatUint(uint64(tag.ID), _buffer)
			_buffer.WriteString(`">`)
			hero.EscapeHTML(tag.Name, _buffer)
			_buffer.WriteString(`</option>
            `)
		}
		_buffer.WriteString(`
          </select>
        </div>
        <div class="col">
            <button type="button" id="reregister-password" data-toggle="modal" data-target="#addModal" class="btn btn-primary">タグを追加する</button>
        </div>
      </div>
      <div class="h2" id="tagname">観察ページ</div>
      <hr>
      <div class="h3 sub">直近 1 週間の様子</div>
      <div class="row">
        <div class="col-xs-12 col-md-8"><div id="week-chart"></div></div>
        <div class="col-xs-12 col-md-4 json" id="week-json"></div>
      </div>
      <div id="week-pagination">
        <button class="prev btn btn-secondary">
          <i class="fa fa-lg fa-chevron-left" aria-hidden="true"></i>
        </button>
        <button class="next btn btn-secondary">
            <i class="fa fa-lg fa-chevron-right" aria-hidden="true"></i>
        </button>
      </div>
      <hr>
      <div class="h3 sub">直近 1 ヶ月の様子</div>
      <div class="row">
        <div class="col-xs-12 col-md-8"><div id="month-chart"></div></div>
        <div class="col-xs-12 col-md-4 json" id="month-json"></div>
      </div>
      <div id="month-pagination">
        <button class="prev btn btn-secondary">
          <i class="fa fa-lg fa-chevron-left" aria-hidden="true"></i>
        </button>
        <button class="next btn btn-secondary">
            <i class="fa fa-lg fa-chevron-right" aria-hidden="true"></i>
        </button>
      </div>
      <hr>
      <div class="h3 sub">これまでの様子</div>
      <div class="row">
        <div class="col-xs-12 col-md-8"><div id="chart"></div></div>
        <div class="col-xs-12 col-md-4 json" id="json"></div>
      </div>
      <div id="all-pagination">
        <button class="prev btn btn-secondary">
          <i class="fa fa-lg fa-chevron-left" aria-hidden="true"></i>
        </button>
        <button class="next btn btn-secondary">
            <i class="fa fa-lg fa-chevron-right" aria-hidden="true"></i>
        </button>
      </div>
    `)
	}
	_buffer.WriteString(`
  </div>
</div>
<!-- Modal -->
<div class="modal fade" id="addModal" tabindex="-1" role="dialog" aria-labelledby="addModalLabel" aria-hidden="true">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="addModalLabel">新規タグの追加</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body">
        <div class="form-group">
          <label for="username" class="form-control-label">タグの名前:</label>
          <input type="text" class="form-control" id="tag_name" placeholder="タグの名前">
        </div>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-dismiss="modal">閉じる</button>
        <button type="button" id="add-tag" class="btn btn-primary">追加する</button>
      </div>
    </div>
  </div>
</div>
`)

	_buffer.WriteString(`
  </main>
  <footer class="footer">
    <p>© `)
	hero.FormatInt(int64(args.Year()), _buffer)
	_buffer.WriteString(` <a class="text-white" href="https://twitter.com/CodeHex">CodeHex</a></p>
  </footer>
  `)
	_buffer.WriteString(`
  <script src="/assets/js/mypage.js"></script>
`)

	_buffer.WriteString(`
</body>
</html>`)
	w.Write(_buffer.Bytes())

}
