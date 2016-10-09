package main

import "fmt"

func hourGlass(set int, s1 []int, s2 []int, s3 []int) int {
	//fmt.Println("Processing set ", set)

	totals1 := 0
	for _, values1 := range s1 {
		totals1 += values1
	}

	totals3 := 0
	for _, values3 := range s3 {
		totals3 += values3
	}

	return totals1 + totals3 + s2[1]

}

func main() {

	var num, sum int
	max := -63
	// Create two-dimensional array.
	numbers := [6][6]int{}

	/* input each array element's */
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scan(&num)
			numbers[i][j] = num
			//fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
		}
	}

	for i := 0; i+2 < 6; i++ {
		for j := 0; j+2 < 6; j++ {
			setone := numbers[i+0][j : j+3]
			//fmt.Println("slice 1 set", i, setone)
			settwo := numbers[i+1][j : j+3]
			//fmt.Println("slice 2 set", i, settwo)
			setthree := numbers[i+2][j : j+3]
			//fmt.Println("slice 3 set ", i, setthree)
			sum = hourGlass(i, setone, settwo, setthree)

			if sum > max {
				max = sum
				//fmt.Println("New max is ", setmax)
			}
		}
	}

	//fmt.Println("Max is ", setmax)
	fmt.Println(max)

}
