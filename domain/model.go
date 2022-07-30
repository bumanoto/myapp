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
	Name string
}

func FindUsers() []User {
	var users []User
	infla.Db.Find(&users)
	return users
}
