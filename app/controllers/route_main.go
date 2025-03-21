package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

// トップページを表示する関数
// 第一引数: レスポンスライター　クライアントに送信するためのインターフェース
// 第二引数: リクエスト
func top(w http.ResponseWriter, r *http.Request) {
	// セッションを取得
	_, err := session(w, r)
	// セッションが存在しない場合、ログインページを表示
	// 存在する場合、ToDoページにリダイレクト
	if err != nil {
		generateHTML(w, "hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

// ログイン認証を行う関数
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	// ユーザーが存在しない場合、ログインページにリダイレクト
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/login", 302)
	}
	//パスワードが一致した場合、クッキーを設定する
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		// クッキーを作成
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		// クッキーをレスポンスに設定
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "index")
	}
}
