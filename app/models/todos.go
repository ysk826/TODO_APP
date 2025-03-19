package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

// ユーザーに紐づくTodoをDBに作成する関数
func (u *User) CreateTodo(content string) (err error) {
	cmd := `INSERT INTO todos (
			content,
			user_id,
			created_at) VALUES (?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// TodoをDBから取得する関数
func GetTodo(id int) (todo Todo, err error) {
	cmd := `SELECT
				id,
				content,
				user_id,
				created_at
			FROM todos WHERE id = ?`

	todo = Todo{}
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err
}

// 全てのTodoをDBから取得する関数
func GetTodos() (todos []Todo, err error) {
	cmd := `SELECT
				id,
				content,
				user_id,
				created_at
			FROM todos`

	// Query()は複数行を返す
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// rows.Next()は次の行があるかどうかを確認する
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	// rowsはデータベースへのコネクションを持っているため、処理が終わったら必ずClose()を呼ぶ
	// リソースの解放を行う
	rows.Close()

	return todos, err
}

// ユーザーに紐づくTodoをDBから取得する関数
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `SELECT
				id,
				content,
				user_id,
				created_at
			FROM todos WHERE user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}
