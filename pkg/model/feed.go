package model

import (
	"github.com/jinzhu/gorm"
)

type Feed struct {
	*gorm.Model
	Name string
	User string
	URL  string
}
