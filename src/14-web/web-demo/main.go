package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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
	structObj := struct {
		Test string
	}{
		Test: "test",
	}
	s, _ := json.Marshal(structObj)
	_, err := w.Write(s)
	if err != nil {
		return
	}
}

func longHandler(w http.ResponseWriter, req *http.Request) {
	structObj := struct {
		Test string
	}{
		Test: "test",
	}
	time.Sleep(3 * time.Minute)
	s, _ := json.Marshal(structObj)
	_, err := w.Write(s)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/long", longHandler)
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
