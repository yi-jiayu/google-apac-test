package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"github.com/yi-jiayu/google-apac-test-2017/codejam"
	"log"
	"os"
)

type RNG struct {
	K    int
	PA   float64
	PB   float64
	PC   float64
	Memo map[string]float64
}

func NewRNG(K, A, B, C int) RNG {
	return RNG{
		K:    K,
		PA:   float64(A) / 100,
		PB:   float64(B) / 100,
		PC:   float64(C) / 100,
		Memo: make(map[string]float64),
	}
}

func (r RNG) Expectation(X, depth int) float64 {
	key := fmt.Sprintf("%d,%d", X, depth)

	if expectation, exists := r.Memo[key]; exists {
		return expectation
	}

	if depth == 1 {
		expectation := r.PA*float64(X&r.K) +
			r.PB*float64(X|r.K) +
			r.PC*float64(X^r.K)

		r.Memo[key] = expectation

		return expectation
	} else {
		expectation := r.PA*r.Expectation(X&r.K, depth-1) +
			r.PB*r.Expectation(X|r.K, depth-1) +
			r.PC*r.Expectation(X^r.K, depth-1)

		r.Memo[key] = expectation

		return expectation
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	T, err := codejam.GetNumTestCases(scanner)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < T; i++ {
		result := scanner.Scan()
		if !result {
			err := scanner.Err()
			if err != nil {
				log.Fatal(errors.Wrap(err, "error while reading test cases"))
			} else {
				log.Fatal(errors.New("unexpected EOF while reading test cases"))
			}
		}

		nxkabc, err := codejam.StringToInts(scanner.Text(), " ")
		if err != nil {
			log.Fatal(err)
		}

		depth := nxkabc[0]
		X := nxkabc[1]

		rng := NewRNG(nxkabc[2], nxkabc[3], nxkabc[4], nxkabc[5])

		expectation := rng.Expectation(X, depth)

		fmt.Printf("Case #%d: %f\n", i + 1, expectation)
	}
}
