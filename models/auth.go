package models

import (
	"gin_blog/config"

	"gorm.io/gorm"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	err := config.DB.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}

type AuthForm struct {
	Username string
	Password string
}

func (a *AuthForm) Check() (bool, error) {
	return CheckAuth(a.Username, a.Password)
}
