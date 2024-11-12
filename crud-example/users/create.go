package users

import "gorm.io/gorm"

func Create(db *gorm.DB, name string) uint {
	user := new(User)
	user.Name = name
	db.Create(user)
	return user.ID
}
