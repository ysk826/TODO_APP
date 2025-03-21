package controllers

import (
	"net/http"
)

// top()はトップページを表示する関数
// 第一引数: レスポンスライター　クライアントに送信するためのインターフェース
// 第二引数: リクエスト
func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "hello", "layout", "public_navbar", "top")
}
