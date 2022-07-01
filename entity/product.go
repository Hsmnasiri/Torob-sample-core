package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"name"`
	Price string `json:"Price"`
	Types []Type `gorm:"many2many:type_products;"`
}

func (p *Product) SaveProduct() (*Product, error) {

	var err error
	err = DB.Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}
