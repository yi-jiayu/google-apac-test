package main

import (
	"testing"
	"fmt"
)

func TestSampleCases(t *testing.T) {
	cases := []string{"ag", "aa", "abcde", "x"}
	answers := []int{4, 1, 108, 1}

	for i, input := range cases {
		fmt.Printf("Input: %s", input)
		output := count(dedup(permute(input)))
		fmt.Printf(", Output: %d, Expected: %d\n", output, answers[i])
	}
}