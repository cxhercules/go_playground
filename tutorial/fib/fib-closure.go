package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first := 0
	second := 0
	result := 0
	return func() int {
		if second == 0 {
			second++
			return 0
		} else {
			if result == 0 {
				result = first + second
				return result
			} else {
				result = first + second
				first = second
				second = result
				return result
			}
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
