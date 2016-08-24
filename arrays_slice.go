package main

import "fmt"

func main() {
	myArray := [3]int{42, 27, 99}

	mySlice := myArray[:]
	fmt.Println(myArray)

	mySlice = append(mySlice, 100)
	fmt.Println(myArray)
	fmt.Println(len(myArray))
	fmt.Println(mySlice)

}
