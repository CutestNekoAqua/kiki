package model

import "github.com/jinzhu/gorm"

type Account struct {
	*gorm.Model
	Name     string
	APIToken string
}
