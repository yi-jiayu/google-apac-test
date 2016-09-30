package main

import (
	"bufio"
	"github.com/yi-jiayu/google-apac-test-2017/codejam"
	"log"
	"os"
	"fmt"
)

const (
	BAND_SIZE = 4
)

func DistinctNumbersAndFrequencies(input []int) ([]int, map[int]int) {
	frequencies := make(map[int]int)

	for _, n := range input {
		if _, exists := frequencies[n]; exists {
			frequencies[n]++
		} else {
			frequencies[n] = 1
		}
	}

	distinct := []int{}

	for k := range frequencies {
		distinct = append(distinct, k)
	}

	return distinct, frequencies
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	T, err := codejam.GetNumTestCases(scanner)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < T; i++ {
		nk, err := codejam.ReadLineOfInts(scanner)
		if err != nil {
			log.Fatal(err)
		}

		K := nk[1]
		var distinctCandidates [4][]int
		var numberFrequencies [4]map[int]int

		for j := 0; j < BAND_SIZE; j++ {
			numbers, err := codejam.ReadLineOfInts(scanner)
			if err != nil {
				log.Fatal(err)
			}

			dC, nF := DistinctNumbersAndFrequencies(numbers)
			distinctCandidates[j] = dC
			numberFrequencies[j] = nF
		}

		// find all distinct combinations of AB and CD
		combAB := make(map[int][][2]int)
		combCD := make(map[int][][2]int)

		for _, A := range distinctCandidates[0] {
			for _, B := range distinctCandidates[1] {
				key := A ^ B
				if val, exists := combAB[key]; !exists {
					combAB[key] = [][2]int{[2]int{A, B}}
				} else {
					combAB[key] = append(val, [2]int{A, B})
				}
			}
		}

		for _, C := range distinctCandidates[2] {
			for _, D := range distinctCandidates[3] {
				key := C ^ D ^ K
				if val, exists := combCD[key]; !exists {
					combCD[key] = [][2]int{[2]int{C, D}}
				} else {
					combCD[key] = append(val, [2]int{C, D})
				}
			}
		}

		// find all the combinations of A^B and C^D^K which are equal to each other
		trendyBands := make([][4]int, 0)

		for key, ABs := range combAB {
			if CDs, exists := combCD[key]; exists {
				for _, AB := range ABs {
					for _, CD := range CDs {
						ABCD := [4]int{AB[0], AB[1], CD[0], CD[1]}
						trendyBands = append(trendyBands, ABCD)
					}
				}
			}
		}

		// count number of possible bands, taking into account repeated robot numbers
		accum := 0

		// for each trendy band, multiply the frequencies of the number in the i-th spot
		// in the set of candidates for that slot together
		for _, band := range trendyBands {
			A, B, C, D := band[0], band[1], band[2], band[3]
			accum += numberFrequencies[0][A] *
				numberFrequencies[1][B] *
				numberFrequencies[2][C] *
				numberFrequencies[3][D]
		}

		fmt.Printf("Case #%d: %d\n", i + 1, accum)
	}
}
