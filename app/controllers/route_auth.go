package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

// signUp()はサインアップページを表示する関数
func signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Panicln(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Panicln(err)
		}
		http.Redirect(w, r, "/", 302)
	}
}

// login()はログインページを表示する関数
func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "login")
}
