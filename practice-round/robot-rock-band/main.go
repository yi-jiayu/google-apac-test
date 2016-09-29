package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"github.com/yi-jiayu/google-apac-test-2017/codejam"
	"log"
	"os"
)

// Steps:
// Extract distinct robot numbers auditioning for each spot
// Note number of robots with same numbers
// Permute all combinations of distinct robot numbers
// For each permutation, find the number of possible bands by multiplying the numbers of robots with each number

const (
	BAND_SIZE = 4
)

type Groups struct {
	Count  int
	Sizes  []int
	Groups [][]int
}

type Permutation struct {
	Groups  *Groups
	Cursors []int
}

func PopulateGroups(groups [][]int) Groups {
	length := len(groups)

	sizes := []int{}
	for _, group := range groups {
		sizes = append(sizes, len(group))
	}

	return Groups{
		length,
		sizes,
		groups,
	}
}

func NewPermutation(groups *Groups) Permutation {
	cursors := []int{}

	for i := 0; i < groups.Count; i++ {
		cursors = append(cursors, 0)
	}

	return Permutation{
		groups,
		cursors,
	}
}

func (p Permutation) Next() Permutation {
	for i := 0; i < p.Groups.Count; i++ {
		if p.Cursors[i]+1 < p.Groups.Sizes[i] {
			p.Cursors[i]++
			break
		} else {
			p.Cursors[i] = 0
		}
	}

	return p
}

func (p Permutation) Values() []int {
	values := []int{}

	for i, cursor := range p.Cursors {
		values = append(values, p.Groups.Groups[i][cursor])
	}

	return values
}

func product(terms ...int) int {
	p := 1
	for _, term := range terms {
		p *= term
	}

	return p
}

func xor(terms ...int) int {
	x := 0
	for _, term := range terms {
		x ^= term
	}

	return x
}

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

	defer func() {
		if err := recover(); err != nil {
			if err != nil {
				log.Fatal(errors.New(fmt.Sprintf("error while processing test cases: %s", err)))
			} else {
				log.Fatal(errors.New("unexpected EOF"))
			}
		}
	}()

	T, err := codejam.GetNumTestCases(scanner)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < T; i++ {
		// get number of robots auditioning
		success := scanner.Scan()
		if !success {
			panic(scanner.Err())
		}

		nk, err := codejam.StringToInts(scanner.Text(), " ")
		if err != nil {
			log.Fatal(errors.New("invalid input"))
		}

		K := nk[1]

		// get candidates
		candidates := [][]int{}

		for j := 0; j < 4; j++ {
			success := scanner.Scan()
			if !success {
				panic(scanner.Err())
			}

			numbers, err := codejam.StringToInts(scanner.Text(), " ")
			if err != nil {
				log.Fatal("invalid input")
			}

			candidates = append(candidates, numbers)
		}

		distinctCandidates := [][]int{}
		candidateFrequencies := []map[int]int{}

		for _, numbers := range candidates {
			d, f := DistinctNumbersAndFrequencies(numbers)
			distinctCandidates = append(distinctCandidates, d)
			candidateFrequencies = append(candidateFrequencies, f)
		}

		groups := PopulateGroups(distinctCandidates)
		perm := NewPermutation(&groups)
		totalPermutations := product(groups.Sizes...)
		validPermutations := [][]int{}

		for l := 0; l < totalPermutations; l++ {
			values := perm.Values()
			xor := xor(values...)
			if xor == K {
				validPermutations = append(validPermutations, values)
			}
			perm = perm.Next()
		}

		// count number of possible bands
		accum := 0

		for _, permutation := range validPermutations {
			count := product(
				candidateFrequencies[0][permutation[0]],
				candidateFrequencies[1][permutation[1]],
				candidateFrequencies[2][permutation[2]],
				candidateFrequencies[3][permutation[3]])

			accum += count
		}

		fmt.Printf("Case #%d: %d\n", i + 1, accum)
	}
}
