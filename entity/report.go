package entity

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	ShopID      uint
}

func (r *Report) SaveReport() (*Report, error) {

	var err error
	err = DB.Create(&r).Error
	if err != nil {
		return &Report{}, err
	}
	return r, nil
}
