package main

import (
	"bufio"
	"fmt"
	"github.com/yi-jiayu/google-apac-test-2017/codejam"
	"log"
	"math"
	"math/rand"
	"os"
)

const (
	EPSILON = 5E-9
)

func value(cashflows []int, r float64) float64 {
	var val float64 = 0
	N := len(cashflows) - 1

	n := N
	v0 := float64(cashflows[0]) * math.Pow(1+r, float64(n))
	val -= v0

	for i := 1; i < N+1; i++ {
		n := N - i
		vn := float64(cashflows[i]) * math.Pow(1+r, float64(n))
		val += vn
	}

	return val
}

func derivative(cashflows []int, r float64) float64 {
	var deriv float64 = 0
	N := len(cashflows) - 1

	// differentiate initial cash outlay
	n := N
	d0 := float64(cashflows[0]*n) * math.Pow(1+r, float64(n-1))
	deriv -= d0

	// differentiate remaining cash inlays
	for i := 1; i < N; i++ {
		n := N - i
		dn := float64(cashflows[i]*n) * math.Pow(1+r, float64(n-1))
		deriv += dn
	}

	return deriv
}

// Calculate the next result using Newton's Method
func IRR(cashflows []int, r float64) float64 {
	return r - value(cashflows, r)/derivative(cashflows, r)
}

// Apply Newton's Method
func iterateIRR(cashflows []int, guess, epsilon float64) float64 {
	var irr0, irr1, eps float64

	irr0 = guess
	irr1 = IRR(cashflows, irr0)
	eps = math.Abs(irr1 - irr0)

	for eps > epsilon {
		irr0 = irr1
		irr1 = IRR(cashflows, irr0)
		eps = math.Abs(irr1 - irr0)
	}

	return irr1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	T, err := codejam.GetNumTestCases(scanner)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < T; i++ {
		_, err := codejam.GetNumTestCases(scanner)
		if err != nil {
			log.Fatal(err)
		}

		cashflows, err := codejam.ReadLineOfInts(scanner)
		if err != nil {
			log.Fatal(err)
		}

		// try with a initial guess of 0
		irr := iterateIRR(cashflows, 0, EPSILON)

		// if irr is not within the correct range, try again with a random initial guess
		for irr > 1 || irr < -1 {
			irr = iterateIRR(cashflows, rand.Float64(), EPSILON)
		}

		fmt.Printf("Case #%d: %.12f\n", i+1, irr)
	}
}
