package main

import (
	"testing"
	"fmt"
)

func TestPermutation_Next(t *testing.T) {
	groups := Groups{
		BAND_SIZE,
		[]int{2, 4, 2},
		[][]int{},
	}

	p := Permutation{
		&groups,
		[]int{0, 0, 0},
	}

	fmt.Println(p.Cursors)

	for i := 0; i < 8; i++ {
		p = p.Next()
		fmt.Println(p.Cursors)
	}
}

func TestDistinctNumbers(t *testing.T) {
	input := []int{1, 2, 2, 3, 3, 4, 5}

	output := DistinctNumbersAndFrequencies(input)

	fmt.Println(output)
}