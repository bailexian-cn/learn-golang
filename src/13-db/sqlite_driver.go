package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main1() {
	db, err := gorm.Open(sqlite.Open("./sqlite_data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var p Product
	db.First(&p, 1)                 // 根据整形主键查找
	db.First(&p, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 Product 的 price 更新为 200
	db.Model(&p).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&p).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&p).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 Product
	//db.Delete(&Product, 1)
}
