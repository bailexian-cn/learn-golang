package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

)

type key int

const (
	KeyTransactionDB key = iota
)

func GetCtxDB(ctx context.Context) *gorm.DB {
	db := ctx.Value(KeyTransactionDB)
	if db == nil {
		return DB.WithContext(ctx)
	}
	if ctxDB, ok := db.(*gorm.DB); ok {
		return ctxDB
	}
	return DB.WithContext(ctx)
}

var DB *gorm.DB

type Config struct {
	Env       string
	EnableLog bool
	DBPath    string
}

func InitDB(dbconfig Config) {
	// verbose db log
	dbLogLever := logger.Warn
	if dbconfig.EnableLog {
		dbLogLever = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      dbLogLever,  // Log level
			Colorful:      true,        // 彩色打印
		},
	)

	var l logger.Interface
	l = newLogger
	_db, err := gorm.Open(mysql.Open(dbconfig.DBPath), &gorm.Config{Logger: l, PrepareStmt: true})
	if err != nil {
		log.Fatal(errors.Wrapf(err, `gorm.Open(mysql.Open(%s), &gorm.Config{})`, dbconfig.DBPath))
	}
	sqldb, err := _db.DB()
	if err != nil {
		log.Fatal(errors.Wrap(err, `_db.DB()`))
	}
	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(100)
	sqldb.SetConnMaxLifetime(time.Second * 600)
	DB = _db
}
