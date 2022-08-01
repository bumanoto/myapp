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

func FindUsersByName(name string) []User {
	var users []User
	infla.Db.Where("name = ?", name).First(&users)
	return users
}

func FindUsersByEmail(email string) []User {
	var users []User
	infla.Db.Where("email = ?", email).First(&users)
	return users
}

func FindUsersByNameOrEmail(nameOrEmail string) []User {
	users := FindUsersByEmail(nameOrEmail)
	if len(users) > 0 {
		return users
	}
	return FindUsersByName(nameOrEmail)
}

func CreateUser(
	name string,
	email string,
	password string) {
	user := &User{
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
