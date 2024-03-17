package models

import "gorm.io/gorm"

type City struct {
	gorm.Model
	Code string
	Name string
}
