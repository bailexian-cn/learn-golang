package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"web-demo/db"
	"web-demo/model"
)

type key int

const (
	Bb key = iota
)

func sleepOneMinute() {
	time.Sleep(20 * 1000 * time.Millisecond)
	fmt.Println("sleep over")
}

func handler(w http.ResponseWriter, req *http.Request) {

	mysqlDb := db.MysqlDb{
		MysqlConfig: db.MysqlConfig{
			Host: "10.255.253.46",
			Port: "3306",
			User: "root",
			Password: "123456",
		},
		Database: "test",
	}
	// 初始化数据库
	dbConf := db.Config{
		Env: "",
		EnableLog: false,
		DBPath: "",
	}

	url := req.URL
	go sleepOneMinute()
	_, _ = fmt.Fprintf(w, "Url of request: %s\n", url)
	_, _ = fmt.Fprintf(w, "Aa=%d, Bb=%d", model.Aa, Bb)
	con := req.Context()
	con = context.WithValue(con, model.Aa, mysqlDb)
	con = context.WithValue(con, Bb, dbConf)
	aa := con.Value(model.Aa)
	fmt.Println(aa)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
