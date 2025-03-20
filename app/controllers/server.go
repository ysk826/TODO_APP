package controllers

import (
	"net/http"
	"todo_app/config"
)

// StattMainServer()はサーバーを起動する関数
func StartMainServer() error {
	// HandleFunc()は指定したパスに対するリクエストを処理するハンドラを登録する
	// 第一引数: パス
	// 第二引数: ハンドラ
	http.HandleFunc("/", top)

	// ListenAndServe()は指定したポートでサーバーを起動する
	// 第一引数: ポート番号
	// 第二引数: ハンドラ
	// 返り値: エラー
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
