package domain

import (
	"gorm.io/gorm"
	"myapp/infla"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type Idea struct {
	gorm.Model
	Name   string
	UserID uint
}

func FindUsers() []User {
	var users []User
	infla.Db.Find(&users)
	return users
}

func CreateUser(
	name string,
	email string,
	password string,
	passwordConfirm string) {
	user := User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	infla.Db.Create(user)
}

func FindIdeas() []Idea {
	var ideas []Idea
	infla.Db.Find(&ideas)
	return ideas
}
