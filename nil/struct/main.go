package main

import (
	"fmt"
)

type myNumber struct {
	n int
}

func plusOne(number *myNumber) {
	number.n++
}

func main() {
	var number *myNumber
	fmt.Printf("%v\n", number)
	plusOne(number)

}
