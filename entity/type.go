package entity

import (
	"gorm.io/gorm"
)

type Type struct {
	gorm.Model
	Name     string `json:"name"`
	Products []Product
}

func GetTypes() ([]Type, error) {
	types := []Type{}
	DB.Find(&types)
	return types, nil
}
func GetTypeProducts(name string) ([]Type, error) {
	var types []Type

	err := DB.Model(&Product{}).Where("name = ?", name).Preload("Products").Find(&types).Error
	return types, err
}
func (t *Type) SaveType() (*Type, error) {

	var err error
	err = DB.Create(&t).Error
	if err != nil {
		return &Type{}, err
	}
	return t, nil
}
