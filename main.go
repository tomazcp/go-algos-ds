package main

import (
	"fmt"
	algo "go-algos-ds/pkg/algos"
)

func main() {
	hs := []int{1, 3, 5, 9, 12, 15, 23, 29, 32, 54}

	fmt.Printf("The needle is at index %d\n", algo.Bsearch(5, hs))
	fmt.Printf("The needle is at index %d\n", algo.Bsearch(29, hs))
}