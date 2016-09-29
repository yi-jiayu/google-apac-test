package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"
	"github.com/pkg/errors"
)

const (
	MODULO_CONSTANT = 1000000007
)

// Given an input word, permute outputs all acceptable letters for each index of the input word as an array of strings.
func permute(input string) []string {
	output := []string{}

	for pos := range input {
		if len(input) == 1 {
			return append(output, input)
		}

		switch pos {
		case 0:
			output = append(output, input[:2])
		case len(input) - 1:
			output = append(output, input[len(input)-2:])
		default:
			output = append(output, input[pos-1:pos+2])
		}
	}

	return output
}

func dedup2(letters string) string {
	if len(letters) == 1 {
		return letters
	}

	if letters[0] == letters[1] {
		return letters[0:1]
	} else {
		return letters
	}
}

func dedup3(letters string) string {
	if letters[0] == letters[1] && letters[1] == letters[2] && letters[0] == letters[2] {
		// three same letters
		return letters[0:1]
	} else if letters[0] != letters[1] && letters[1] != letters[2] && letters[0] != letters[2] {
		// three different letters
		return letters
	} else if letters[0] == letters[1] {
		// indexes 1 and 2 are distinct
		return letters[1:]
	} else {
		// indexes 0 and 1 are distinct
		return letters[:2]
	}
}

// Given the output from permute, dedup removes repeated letters from each index
func dedup(input []string) []string {
	output := []string{}

	for pos, letters := range input {
		switch pos {
		case 0:
			output = append(output, dedup2(letters))
		case len(input) - 1:
			output = append(output, dedup2(letters))
		default:
			output = append(output, dedup3(letters))
		}
	}

	return output
}

func count(input []string) int {
	count := 1

	for _, letters := range input {
		count = count * len(letters)
		count = count % MODULO_CONSTANT
	}

	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// scan first line (number of test cases)
	success := scanner.Scan()
	if !success {
		err := scanner.Err()
		if err != nil {
			log.Fatal(errors.Wrap(err, "err: error while reading number of test cases"))
		} else {
			log.Fatal("err: unexpected EOF before reading number of test cases")
		}
	}
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(errors.Wrap(err, "err: failed to parse number of test cases"))
	}

	// run the next n test cases
	for i := 0; i < n; i++ {
		success := scanner.Scan()
		if !success {
			err := scanner.Err()
			if err != nil {
				log.Fatal(errors.Wrap(err, "err: error while processing test cases"))
			} else {
				log.Fatal("err: unexpected EOF before expected number of test cases")
			}
		}
		input := scanner.Text()
		answer := count(dedup(permute(input)))
		fmt.Printf("Case #%d: %d\n", i + 1, answer)
	}
}
