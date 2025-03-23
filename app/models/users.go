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
	Todos     []Todo
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
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

// emailからユーザーを取得する関数
func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `SELECT
				id,
				uuid,
				name,
				email,
				password,
				created_at
			FROM users WHERE email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)
	return user, err
}

// ユーザーのセッションを作成する関数
// セッションはログイン状態を保持するためのもの
// セッションはログイン時に作成され、ログアウト時に削除される
// セッションはユーザーごとに一つだけ作成される
// セッションはユーザーがログインしているかどうかを判定するために使われる
func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `INSERT INTO sessions (
		uuid,
		email,
		user_id,
		created_at) VALUES (?, ?, ?, ?)`
	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	cmd2 := `SELECT
			id,
			uuid,
			email,
			user_id,
			created_at
		FROM sessions WHERE user_id = ? AND email = ?`
	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt)

	return session, err
}

// セッションが存在するかチェックする関数
func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `SELECT
				id,
				uuid,
				email,
				user_id,
				created_at
			FROM sessions WHERE uuid = ?`
	// QueryRow()は一行を返す
	// Scan()はQueryRow()の返り値を引数に取り、引数に指定した変数に値をセットする
	// クエリの結果がない、型の不一致の場合はerrに値がセットされる
	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	// IDが0ならセッションが存在しない
	// IDは通常1から始まるため
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

// セッションをUUIDで削除する関数
func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `DELETE FROM sessions WHERE uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// セッションからユーザーを取得する関数
func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `SELECT
				id,
				uuid,
				name,
				email,
				created_at
			FROM users WHERE id = ?`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt)
	return user, err
}
