package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string `json:"name"`
	LowestPrice  string `json:"lowest_price"`
	HighestPrice string `json:"highest_price"`
	TypeID       uint
	SubTypeID    uint
	Img          string
	Shops        []Shop `gorm:"many2many:shop_products;"`
}
type shop_products struct {
	gorm.Model
	shop_id    uint
	product_id uint
}

func IncrementShop(sid uint, pid uint) error {
	uf := new(shop_products)
	uf.shop_id = sid
	uf.product_id = pid
	err := DB.Create(uf).Error
	return err
}
func GetProducts() ([]Product, error) {
	products := []Product{}
	DB.Find(&products)
	return products, nil
}
func GetOneProduct() ([]Product, error) {
	products := []Product{}
	DB.Find(&products)
	return products, nil
}
func (p *Product) SaveProduct() (*Product, error) {

	err := DB.Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}
