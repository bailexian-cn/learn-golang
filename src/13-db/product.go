package main

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string `gorm: "code"`
	Price uint   `gorm: "price"`
}
