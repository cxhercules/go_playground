package main

import "fmt"

func main() {
	var num_in int
	fmt.Scan(&num_in)

	numbers := make([]int, num_in)
	var input int

	for i := 0; i < num_in; i++ {
		fmt.Scan(&input)
		numbers[i] = input
	}

	bubbleSort(numbers)
	fmt.Printf("First Element: %d\n", numbers[0])
	fmt.Printf("Last Element: %d\n", numbers[len(numbers)-1])
}

func bubbleSort(numbers []int) {
	var N int = len(numbers)
	var i int
	numSwaps := 0
	for i = 0; i < N; i++ {
		if !sweep(numbers, i, &numSwaps) {
			return
		}
	}
}

func sweep(numbers []int, prevPasses int, Swaps *int) bool {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool = false

	for secondIndex < (N - prevPasses) {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
			*Swaps++
		}

		firstIndex++
		secondIndex++
	}

	if !didSwap {
		fmt.Printf("Array is sorted in %d swaps.\n", *Swaps)
	}
	return didSwap
}
