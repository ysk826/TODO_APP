package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo_app/app/models"
	"todo_app/config"
)

// generateHTMLはHTMLを生成する関数
// 第一引数: レスポンスライター
// 第二引数: テンプレートに渡すデータ
// 第三引数: ファイル名
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	// Must()はエラーがあればpanicを起こす
	// ParseFiles()は指定したファイルを読み込む
	// ExecuteTemplate()は読み込んだファイルを実行する
	// 第一引数: レスポンスライター
	// 第二引数: {{define "layout"}}{{end}}のように名前をつけたテンプレートを渡す
	// 第三引数: ファイル名
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// セッションを取得する関数
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	// クッキーが存在する場合、セッションを取得
	if err == nil {
		// クッキーの値をセッションのUUIDとしてセット
		sess = models.Session{UUID: cookie.Value}
		// sessがDBに存在するか確認
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

// サーバーを起動する関数
func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// HandleFunc()は指定したパスに対するリクエストを処理するハンドラを登録する
	// 第一引数: パス
	// 第二引数: ハンドラ
	// 第一引数のパスにリクエストが来た場合、第二引数のハンドラが呼び出される
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)

	// ListenAndServe()は指定したポートでサーバーを起動する
	// 第一引数: ポート番号
	// 第二引数: ハンドラ
	// 返り値: エラー
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
