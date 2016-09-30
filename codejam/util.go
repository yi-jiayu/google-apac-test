package codejam

import (
	"strconv"
	"github.com/pkg/errors"
	"strings"
)

type Scanner interface {
	Scan() bool
	Err() error
	Text() string
}

func GetNumTestCases(scanner Scanner) (int, error) {
	success := scanner.Scan()
	if !success {
		err := scanner.Err()
		if err != nil {
			return 0, errors.Wrap(err, "err: error while reading number of test cases")
		} else {
			return 0, errors.New("err: unexpected EOF before reading number of test cases")
		}
	}
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, errors.Wrap(err, "err: failed to parse number of test cases")
	}

	return n, nil
}

// Takes an input string of delimited integers and returns an array of integers.
func StringToInts(s string, sep string) ([]int, error) {
	sInts := strings.Split(s, sep)
	ints := []int{}

	for _, sInt := range sInts {
		int_, err := strconv.Atoi(sInt)
		if err != nil {
			return nil, errors.New("invalid input")
		}

		ints = append(ints, int_)
	}

	return ints, nil
}

func ReadLineOfInts(scanner Scanner) ([]int, error) {
	success := scanner.Scan()
	if !success {
		err := scanner.Err()
		if err != nil {
			return nil, err
		} else {
			return nil, errors.New("unexpected EOF")
		}
	}

	ints, err := StringToInts(scanner.Text(), " ")
	if err != nil {
		return nil, err
	}

	return ints, nil
}

