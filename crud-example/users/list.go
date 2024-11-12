package users

import "gorm.io/gorm"

func List(db *gorm.DB) []User {
	var users []User
	db.Find(&users)
	return users
}
