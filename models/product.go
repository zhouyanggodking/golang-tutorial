package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Number string
	UserID uint
}