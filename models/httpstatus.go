package models

import (
	"gorm.io/gorm"
)

type Codes struct {
	HttpStatusCodes map[string]HttpStatusCode `yaml:"httpstatuscodes"`
}

type HttpStatusCode struct {
	*gorm.Model
	ID      uint   `gorm:"autoIncrement"`
	Code    int    `gorm:"index;unique"`
	Name    string `gorm:"not null"`
	Desc    string `gorm:"not null"`
	Picture string `gotm:"not null"`
}
