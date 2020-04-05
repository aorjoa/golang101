package model

import "github.com/jinzhu/gorm"

// Product is a product
type Product struct {
	gorm.Model
	Code  string
	Price uint
}
