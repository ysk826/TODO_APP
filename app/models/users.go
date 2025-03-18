package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// ユーザーをDBに作成する関数
func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at) VALUES (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// ユーザーをDBから取得する関数
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `SELECT
				id,
				uuid,
				name,
				email,
				password,
				created_at
			FROM users WHERE id = ?`
	// QueryRow()は一行を返す
	// Scan()はQueryRow()の返り値を引数に取り、引数に指定した変数に値をセットする
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)
	return user, err
}

// ユーザーを更新する関数
func (u *User) UpdateUser() (err error) {
	cmd := `UPDATE users SET
			name = ?,
			email = ?
			WHERE id = ?`
	// Exec()はコマンドの結果を返すが、今回は使わないので_で受ける
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// ユーザーを削除する関数
func (u *User) DeleteUser() (err error) {
	cmd := `DELETE FROM users WHERE id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
