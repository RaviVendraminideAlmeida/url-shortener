package models

import (
	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	Original string `json:"original" gorm:"text; not null; default: ''"`
	Short    string `json:"short" gorm:"text; not null; default: ''"`
}
