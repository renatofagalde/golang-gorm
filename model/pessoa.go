package model

import "gorm.io/gorm"

type Pessoa struct {
	gorm.Model
	Nome        string
	Idade       int8
	PhoneNumber string
}
