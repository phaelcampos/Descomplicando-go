package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/phaelcampos/crud-example/users"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("meubanco.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&users.User{})
	service := users.UserService{DB: db}

	id := service.Create("Raphael")
	service.Update(id, "joão")
	userList := service.List()
	for _, user := range userList {
		fmt.Println(user.Name)
	}
	service.Delete(id)
	service.List()
}
