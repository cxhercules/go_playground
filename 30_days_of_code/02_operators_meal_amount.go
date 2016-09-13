package main

import "fmt"

func main() {

	// Declare second integer, double, and String variables.
	var mealCost float32
	var tipPercent float32
	var taxPercent float32
	var tip float32
	var tax float32
	var totalCost float32

	// Read and save an integer, double, and String to your variables.
	fmt.Scan(&mealCost)
	fmt.Scan(&tipPercent)
	fmt.Scan(&taxPercent)

	// Print the sum of both integer variables on a new line.
	//fmt.Printf("%.1f\t%.1f\t%.1f\n", mealCost, tipPercent, taxPercent)

	// Calculate tip
	tip = mealCost * (tipPercent / 100)

	// Calculate Tax Amount
	tax = mealCost * (taxPercent / 100)

	// Total Cost
	totalCost = mealCost + tip + tax

	fmt.Printf("The total meal cost is %.0f dollars.\n", totalCost)
}
