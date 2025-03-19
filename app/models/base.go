package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
)

func init() {
	// Open()はデータベースへの接続を確立する
	// 第一引数: ドライバ名（データベースの種類）
	// 第二引数: データベース名
	// 返り値: データベース接続のインスタンス
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	// %sはtableNameUserに置き換えられる
	// CREATE TABLE IF NOT EXISTSはテーブルが存在しない場合にテーブルを作成する
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)

	// %sはtableNameTodoに置き換えられる
	// CREATE TABLE IF NOT EXISTSはテーブルが存在しない場合にテーブルを作成する
	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	Db.Exec(cmdT)
}

// ユーザーを作成する関数
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// パスワードをハッシュ化する関数
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
