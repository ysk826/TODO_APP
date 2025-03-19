package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	/*
		u := &models.User{}
		u.Name = "test"
		u.Email = "example.com"
		u.Password = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/
	/*
		u, _ := models.GetUser(2)
		fmt.Println(u)
	*/
	/*
		u.Name = "Test3"
		u.Email = "test3@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	/*
		u.CreateTodo("First Todo")
	*/

	t, _ := models.GetTodo(1)
	fmt.Println(t)
}
