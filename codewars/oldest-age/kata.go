package main

import "fmt"

func TwoOldestAges(ages []int) [2]int {
	twoOld := [2]int{0, 0}

	for i := 0; i < len(ages); i++ {

		if ages[i] > twoOld[1] {
			twoOld[0] = twoOld[1]
			twoOld[1] = ages[i]
		} else if ages[i] > twoOld[0] {
			twoOld[0] = ages[i]
		}

	}
	return twoOld
}

func main() {
	s := [2]int{}
	s = TwoOldestAges([]int{1, 5, 87, 45, 8, 8})

	fmt.Println(s)
}
