package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string `json:"name"`
	LowestPrice  string `json:"lowest_price"`
	HighestPrice string `json:"highest_price"`
	TypeID       uint
	SubTypeID    uint
	Shops        []Shop `gorm:"many2many:shop_products;"`
}

func (p *Product) SaveProduct() (*Product, error) {

	var err error
	err = DB.Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}
