package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

type MysqlDb struct {
	MysqlConfig
	Database string
}

func (db *MysqlDb) getPath() string {
	return db.User + ":" + db.Password +
		"@tcp(" + db.Host + ":" + db.Port + ")/" + db.Database +
		"?charset=utf8mb4&parseTime=True&loc=Local"
}

func main() {
	localDb := MysqlDb{
		MysqlConfig: MysqlConfig{
			Host:     "localhost",
			Port:     "3306",
			User:     "root",
			Password: "123456",
		},
		Database: "go_db-demo",
	}
	db, err := gorm.Open(mysql.Open(localDb.getPath()), &gorm.Config{})

	if err != nil {
		panic("failed to connect mysql database")
	}

	db.AutoMigrate(&Product{})
	//newRow := Product{
	//	Code: "XYJ04",
	//	Price: 400}
	// create
	//db.Create(&newRow)
	// select
	var product Product
	db.First(&product, 2)
	//fmt.Println(product)
	//db.First(&product, "code = ?", "XYJ04")
	//fmt.Println(product)
	//db.Model(&product).Updates(Product{Price: 300, Code: "XYJ44"})
	db.Delete(&product)
}
