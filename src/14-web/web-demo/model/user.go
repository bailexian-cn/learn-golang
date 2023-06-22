package model

type key int

const (
	Aa key = iota
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}
