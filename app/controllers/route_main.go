package controllers

import (
	"html/template"
	"log"
	"net/http"
)

// top()はトップページを表示する関数
// 第一引数: レスポンスライター　クライアントに送信するためのインターフェース
// 第二引数: リクエスト
func top(w http.ResponseWriter, r *http.Request) {
	// ParseFiles()は指定したファイルを読み込む
	// Execute()は読み込んだファイルを実行する
	// 第一引数: レスポンスライター
	// 第二引数: テンプレートに渡すデータ
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, "nil")
}
