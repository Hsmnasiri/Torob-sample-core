package entity

import (
	"gorm.io/gorm"
)

type SubType struct {
	gorm.Model
	Name     string `json:"name"`
	Products []Product
}

func (t *SubType) SaveSubType() (*SubType, error) {

	var err error
	err = DB.Create(&t).Error
	if err != nil {
		return &SubType{}, err
	}
	return t, nil
}
