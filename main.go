package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	fmt.Println(models.Db)

	/*
		u := &models.User{}
		u.Name = "test"
		u.Email = "example.com"
		u.Password = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/

	u, _ := models.GetUser(1)
	fmt.Println(u)

	u.Name = "Test3"
	u.Email = "test3@example.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

}
