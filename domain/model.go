package domain

import (
	"gorm.io/gorm"
	"myapp/infla"
)

type User struct {
	gorm.Model
	Name string
}

type Idea struct {
	gorm.Model
	Name   string
	UserID uint
	User   `gorm:"foreignKey:UserID"`
}

func FindUsers() []User {
	var users []User
	infla.Db.Find(&users)
	return users
}

func FindIdeas() []Idea {
	var ideas []Idea
	infla.Db.Find(&ideas)
	return ideas
}
