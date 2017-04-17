package main

import (
	"fmt"
)

// EvenOrOdd check even or odd
func EvenOrOdd(number int) string {

	if number%2 == 0 {
		return "Even"
	}
	return "Odd"

}

func main() {
	fmt.Println(EvenOrOdd(3))
}
