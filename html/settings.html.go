// Code generated by hero.
// source: /Users/codehex/Desktop/go/src/github.com/Code-Hex/vegeta/template/settings.html
// DO NOT EDIT!
package html

import (
	"io"

	"github.com/shiyanhui/hero"
)

func Settings(args SettingsArgs, w io.Writer) {
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
  <script src="/assets/js/main.js"></script>
`)

	_buffer.WriteString(`
  <title>`)
	_buffer.WriteString(`settings`)

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
              `)
		if args.IsAdmin() {
			_buffer.WriteString(`
                <a class="dropdown-item" href="/mypage/admin"><i class="fa fa-lock" aria-hidden="true"></i> ユーザー管理パネル</a>
              `)
		} else {
			_buffer.WriteString(`
                <a class="dropdown-item" href="/mypage/settings"><i class="fa fa-cog" aria-hidden="true"></i> 設定</a>
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

	settingsArgs := args
	user := settingsArgs.User()

	_buffer.WriteString(`
<input type="hidden" id="api-token" value="`)
	hero.EscapeHTML(settingsArgs.Token(), _buffer)
	_buffer.WriteString(`">
<div class="app-details">
  <div class="container">
    <div class="row">
      <div class="col-xs-12 col-md-6">
        <h3>アクセストークンの変更</h3>
        <div class="form-group">
          <label for="access-token">アクセストークン</label>
          <input type="text" class="form-control" id="access-token" value="`)
	_buffer.WriteString(user.Token)
	_buffer.WriteString(`" readonly>
        </div>
        <button type="button" id="regen-token" class="btn btn-primary btn-lg float-right">アクセストークンを更新する</button>
      </div>
    </div>
  </div>
</div>
<div class="app-details">
  <div class="container">
    <div class="row">
      <div class="col-xs-12 col-md-6">
        <h3>パスワードの変更</h3>
        <div class="form-group">
          <label for="password">パスワード</label>
          <input type="password" class="form-control" id="password" required>
        </div>
        <div class="form-group">
          <label for="password-verify">パスワードの再確認</label>
          <input type="password" class="form-control" id="password-verify" required>
        </div>
        <button type="button" id="reregister-password" class="btn btn-primary btn-lg float-right">パスワードを変更する</button>
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
  <script src="/assets/js/settings.js"></script>
`)

	_buffer.WriteString(`
</body>
</html>`)
	w.Write(_buffer.Bytes())

}
