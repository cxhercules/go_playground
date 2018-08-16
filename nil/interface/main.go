package main

import (
	"errors"
	"fmt"
	"log"
)

type Named interface {
	Name() string
}

func greeting(thing Named) (string, error) {
	if thing == nil {
		return "", errors.New("thing cannot be nil")
	}
	return "Hello " + thing.Name(), nil
}

func main() {
	fmt.Println("vim-go")
	str, err := greeting(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
