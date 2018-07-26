package main

import "fmt"

type Named interface {
	Name() string
}

func greeting(thing Named) string {
	return "Hello " + thing.Name()
}

func main() {
	fmt.Println("vim-go")
	greeting(nil)
}
