package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Test struct {
	Code int    `json:"code,omitempty"`
	Data string `json:"data"`
}

func TestName(t *testing.T) {
	t1 := Test{
		Code: 0,
		Data: "hah",
	}
	byteArr, err := json.Marshal(t1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(byteArr))
}
