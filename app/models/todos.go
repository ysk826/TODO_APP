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
