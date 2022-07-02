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
func (t *Type) SaveType() (*Type, error) {

	var err error
	err = DB.Create(&t).Error
	if err != nil {
		return &Type{}, err
	}
	return t, nil
}
