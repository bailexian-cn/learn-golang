package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func demoHandler(w http.ResponseWriter, req *http.Request) {
	structObj := struct {
		Demo string
	}{
		Demo: "demo",
	}
	s, _ := json.Marshal(structObj)
	_, err := w.Write(s)
	if err != nil {
		return
	}
}

func demoServer() {
	http.HandleFunc("/", demoHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func main() {
	demoServer()
}
