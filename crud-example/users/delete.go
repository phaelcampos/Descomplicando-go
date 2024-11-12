package users

import "gorm.io/gorm"

func Delete(db *gorm.DB, id uint) {
	db.Delete(&User{}, id)
}
