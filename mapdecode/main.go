package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Firstname string
	Lastname  string
	Address   struct {
		City  string
		State string
	}
}

func main() {
	mapPerson := make(map[string]interface{})
	var person Person
	mapPerson["firstname"] = "Nic"
	mapPerson["lastname"] = "Raboy"
	mapAddress := make(map[string]interface{})
	mapAddress["city"] = "San Francisco"
	mapAddress["state"] = "California"
	mapPerson["address"] = mapAddress
	mapstructure.Decode(mapPerson, &person)
	fmt.Println(person)
}
