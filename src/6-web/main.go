package main

import (
	"bailexian.cn/learn-golang/6-web/util"
	"fmt"
)

func main()  {
	if mask, err := util.GetMaskStrFromCIDR("192.168.122.0/24"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("mask is %s\n", mask)
	}
}
