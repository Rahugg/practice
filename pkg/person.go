package pkg

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Id      uint64 `json: "id"`
	Name    string `json : "name"`
	Surname string `json: "surname"`
	Age     int    `json: "age"`
	Address string `json: "address"`
}
