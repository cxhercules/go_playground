package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input uint64
	var count, max int
	fmt.Scan(&input)

	binary_string := strconv.FormatUint(input, 2)
	//fmt.Println(binary_string)

	for _, r := range binary_string {
		c := string(r)
		//fmt.Println(c)
		if c == "1" {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}
	fmt.Println(max)

}
