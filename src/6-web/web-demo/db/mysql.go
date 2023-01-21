package db

type MysqlConfig struct {
	Host string
	Port string
	User string
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
