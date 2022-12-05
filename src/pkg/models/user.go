package models

import (
	"github.com/DanielT777/Web-Services-Shad/src/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB = config.GetDB()

type User struct {
	gorm.Model
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	PseudoName string `json:"pseudo_name"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
	Password   string `json:"password"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserByID(id int) *User {
	var user User
	db.First(&user, id)
	return &user
}

func (u *User) UpdateUser() *User {
	db.Save(&u)
	return u
}

func DeleteUser(id int) {
	var user User
	db.First(&user, id)
	db.Delete(&user)
}
