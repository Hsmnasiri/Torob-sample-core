package entity

import (
	"gorm.io/gorm"
)

type SubType struct {
	gorm.Model
	Name     string `json:"name"`
	Products []Product
}

func GetSubTypes() ([]Type, error) {
	types := []Type{}
	DB.Find(&types)
	return types, nil
}
func GetSubTypeProducts(name string) ([]Type, error) {
	var types []Type

	err := DB.Model(&Product{}).Where("name = ?", name).Preload("Products").Find(&types).Error
	return types, err
}
func (t *SubType) SaveSubType() (*SubType, error) {

	var err error
	err = DB.Create(&t).Error
	if err != nil {
		return &SubType{}, err
	}
	return t, nil
}
