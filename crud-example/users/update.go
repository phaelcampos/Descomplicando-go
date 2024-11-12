package users

import "gorm.io/gorm"

func Update(db *gorm.DB, id uint, name string) {
	var user User
	db.First(&user, id)
	db.Model(&user).Update("Name", name)
}
