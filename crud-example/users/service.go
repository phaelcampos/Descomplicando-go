package users

import "gorm.io/gorm"

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) List() []User {
	var users []User
	s.DB.Find(&users)
	return users
}

func (s *UserService) Delete(id uint) {
	s.DB.Delete(&User{}, id)
}
func (s *UserService) Create(name string) uint {
	user := new(User)
	user.Name = name
	s.DB.Create(user)
	return user.ID
}
func (s *UserService) Update(id uint, name string) {
	var user User
	s.DB.First(&user, id)
	s.DB.Model(&user).Update("Name", name)
}
