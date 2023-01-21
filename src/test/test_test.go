package main

import (
	"fmt"
	"net"
	"testing"
)

type Cat interface {
	Meow()
}

type Tabby struct {}
func (*Tabby) Meow() { fmt.Println("meow") }

func GetACat() Cat {
	var myTabby *Tabby = nil
	// Oops, we forgot to set myTabby to a real value
	return myTabby
}

func TestGetACat(t *testing.T) {
	ip, ipnet, _ := net.ParseCIDR("192.168.1.1/24")
	fmt.Println(ip)
	fmt.Println(ipnet.IP)
}
