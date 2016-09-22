package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	mySlice := make([]string, T)
	var input string

	fmt.Println(mySlice)

	for i := 0; i < T; i++ {
		fmt.Scan(&input)
		mySlice[i] = input
	}

	fmt.Println(mySlice)

	for i := len(mySlice) - 1; i >= 0; i-- {
		fmt.Printf("%s ", mySlice[i])
	}

}
