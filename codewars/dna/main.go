package main

import "fmt"

// ns : slice of indices
// xs, ys : chromosomes as slices of ints
func Crossover(ns []int, xs []int, ys []int) ([]int, []int) {
	// Encountered Map
	encountered := map[int]bool{}
	holdSlice := make([]int, len(xs), (cap(xs)))

	for v := range ns {
		if encountered[ns[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[ns[v]] = true
			// Do Swap
			copy(holdSlice[ns[v]:len(xs)], xs[ns[v]:len(xs)])
			copy(xs[ns[v]:len(xs)], ys[ns[v]:len(xs)])
			copy(ys[ns[v]:len(xs)], holdSlice[ns[v]:len(xs)])
		}

	}
	return xs, ys
}

func main() {
	xs, ys := Crossover([]int{1}, []int{1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2})
	fmt.Println(xs)
	fmt.Println(ys)
}
