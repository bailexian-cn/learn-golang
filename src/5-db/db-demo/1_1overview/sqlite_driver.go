package main

import (
	"db-demo/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)



func main1() {
	db, err := gorm.Open(sqlite.Open("./sqlite_data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&model.Product{})

	// Create
	db.Create(&model.Product{Code: "D42", Price: 100})

	// Read
	var Product model.Product
	db.First(&Product, 1) // 根据整形主键查找
	db.First(&Product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 model.Product 的 price 更新为 200
	db.Model(&Product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&Product).Updates(model.Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&Product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 model.Product
	//db.Delete(&model.Product, 1)
}