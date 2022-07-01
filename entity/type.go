package entity

import (
	"gorm.io/gorm"
)

type Type struct {
	gorm.Model
	Name string `json:"name"`
}

func GetTypes() ([]Type, error) {
	types := []Type{}
	DB.Find(&types)
	return types, nil
}
