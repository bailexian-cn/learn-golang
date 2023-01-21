package db

func init() {
	mysqlDb := MysqlDb{
		MysqlConfig: MysqlConfig{
			Host: "10.255.253.46",
			Port: "3306",
			User: "root",
			Password: "123456",
		},
		Database: "test",
	}
	// 初始化数据库
	dbConf := Config{
		Env: "",
		EnableLog: false,
		DBPath: mysqlDb.getPath(),
	}
	InitDB(dbConf)
}
