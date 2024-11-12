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
	id := users.Create(db, "Raphael")
	users.Update(db, id, "joão")
	userList := users.List(db)
	for _, user := range userList {
		fmt.Println(user.Name)
	}
	users.Delete(db, id)
	users.List(db)
}
