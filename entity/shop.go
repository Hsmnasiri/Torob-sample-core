package entity

import (
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Products    []Product `gorm:"many2many:shop_products;"`
	Report      []Report
}

func (s *Shop) SaveShop() (*Shop, error) {

	var err error
	err = DB.Create(&s).Error
	if err != nil {
		return &Shop{}, err
	}
	return s, nil
}
