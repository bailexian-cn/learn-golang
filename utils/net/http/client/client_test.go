package client

import (
	"encoding/json"
	"k8s.io/apimachinery/pkg/util/wait"
	"log"
	"net/http"
	"testing"
	"time"
)

func fastHandler(w http.ResponseWriter, req *http.Request) {
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

func slowHandler(w http.ResponseWriter, req *http.Request) {
	structObj := struct {
		Test string
	}{
		Test: "test",
	}
	time.Sleep(10 * time.Minute)
	s, _ := json.Marshal(structObj)
	_, err := w.Write(s)
	if err != nil {
		return
	}
}

func demoServer() {
	http.HandleFunc("/fast", fastHandler)
	http.HandleFunc("/slow", slowHandler)
	log.Fatal(http.ListenAndServe("127.0.0.1:7777", nil))
}

func TestHttpTimeout(t *testing.T) {
	go demoServer()
	res := map[string]interface{}{}
	e := wait.PollImmediate(500*time.Millisecond, 5*time.Second, func() (done bool, err error) {
		e1 := HttpCall(http.MethodGet, "http://127.0.0.1:7777/fast", nil, nil, nil, &res)
		if _, ok := res["Test"]; ok && e1 == nil {
			return true, nil
		} else {
			return false, nil
		}
	})
	if e != nil {
		t.Error(e)
	}
	HttpClientTimeout = 5 * time.Second
	err := HttpCall(http.MethodGet, "http://127.0.0.1:7777/slow", nil, nil, nil, &res)
	errType := JudgeErrType(err)
	if errType != TypeRequestTimeout {
		t.Error(err)
	}
}
