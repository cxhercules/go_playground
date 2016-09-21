package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	mySlice := make([]string, T)
	var input, even, odd string

	for i := 0; i < T; i++ {
		fmt.Scan(&input)
		mySlice[i] = input
	}

	for x := range mySlice {
		for i := 0; i < len(mySlice[x]); i++ {
			if i%2 == 0 {
				even += string(mySlice[x][i])
			} else {
				odd += string(mySlice[x][i])
			}
		}
		fmt.Printf("%s %s\n", even, odd)
		even, odd = "", ""
	}
}
