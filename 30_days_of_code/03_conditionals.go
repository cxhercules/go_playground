package main

import "fmt"

func main() {

	// Declare single integer variable.
	var num int

	// Read and save an integer.
	fmt.Scan(&num)
	if num < 1 || num > 100 {
		fmt.Println("Number outside accepted value")
	} else if num%2 == 0 && num >= 2 && num <= 5 {
		fmt.Println("Not Weird")
	} else if num%2 == 0 && num >= 6 && num <= 20 {
		fmt.Println("Weird")
	} else if num%2 == 0 && num > 20 {
		fmt.Println("Not Weird")
	} else {
		fmt.Println("Weird")
	}

}
