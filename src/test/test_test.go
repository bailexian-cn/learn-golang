package main

import (
	"fmt"
	"strconv"
	"testing"
)

type Test struct {
	Code int    `json:"code,omitempty"`
	Data string `json:"data"`
}

func TestName(t *testing.T) {
	dpt := int64(7105108852927791000)
	str := strconv.FormatInt(dpt, 10)
	fmt.Println(str)
}
