package main

import (
	"fmt"
	algo "go-algos-ds/pkg/algos"
)

func main() {
	hs := []int{1, 3, 5, 9, 12, 15, 23, 29, 32, 54}

	fmt.Printf("The needle is at index %d\n", algo.Bsearch(5, hs))
	fmt.Printf("The needle is at index %d\n", algo.Bsearch(29, hs))
	fmt.Printf("The needle is at index %d\n", algo.Bsearch(1, hs))
	fmt.Printf("The needle is at index %d\n", algo.Bsearch(54, hs))

	fmt.Printf("The Crystal Ball optimized break is at level: %d\n", algo.TwoCb([]bool{false, false, false, false, false, false, false, false, true, true, true, true}))
}
