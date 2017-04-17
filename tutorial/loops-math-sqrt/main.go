package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const delta = 0.0000000001
	prev, z := 1.0, 1.0
	steps := 0
	for {
		z -= (z*z - x) / (2*z)
		fmt.Printf("x(%d): %f\n", steps, z )
		if math.Abs(prev-z) < delta {
			break
		}
		prev = z
		steps++
	}
	fmt.Println("# steps:", steps)
	return z
}

func main() {
	fmt.Println(Sqrt(81))
}
