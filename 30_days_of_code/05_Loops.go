package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var num int

	fmt.Scan(&num)

	if num >= 2 && num <= 20 {
		for i := 1; i < 11; i++ {
			fmt.Printf("%d x %d = %d\n", num, i, num*i)
		}
	} else {
		fmt.Println("Input not valid, please input number between 2 and 20 inclusively.")
	}
}
