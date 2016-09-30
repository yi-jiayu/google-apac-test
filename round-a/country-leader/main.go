package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"github.com/yi-jiayu/google-apac-test-2017/codejam"
	"log"
	"os"
	"sort"
	"unicode"
)

type Citizen struct {
	Name             string
	NumUniqueLetters int
}

type By func(c1, c2 *Citizen) bool

func (by By) Sort(citizens []Citizen) {
	sorter := &citizenSorter{
		citizens,
		by,
	}
	sort.Sort(sorter)
}

type citizenSorter struct {
	citizens []Citizen
	by       func(c1, c2 *Citizen) bool
}

func (s *citizenSorter) Len() int {
	return len(s.citizens)
}

func (s *citizenSorter) Swap(i, j int) {
	s.citizens[i], s.citizens[j] = s.citizens[j], s.citizens[i]
}

func (s *citizenSorter) Less(i, j int) bool {
	return s.by(&s.citizens[i], &s.citizens[j])
}

func NewCitizen(name string) Citizen {
	letters := make(map[rune]struct{})

	for _, char := range name {
		if unicode.IsUpper(char) {
			letters[char] = struct{}{}
		}
	}

	return Citizen{
		Name:             name,
		NumUniqueLetters: len(letters),
	}
}

func uniqueLettersDesc(c1, c2 *Citizen) bool {
	return c1.NumUniqueLetters > c2.NumUniqueLetters
}

func alphabetical(c1, c2 *Citizen) bool {
	return c1.Name < c2.Name
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	T, err := codejam.GetNumTestCases(scanner)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < T; i++ {
		N, err := codejam.GetNumTestCases(scanner)
		if err != nil {
			log.Fatal(err)
		}

		var citizens []Citizen

		for j := 0; j < N; j++ {
			success := scanner.Scan()
			if !success {
				err := scanner.Err()
				if err != nil {
					log.Fatal(errors.Wrap(err, "error while reading input"))
				} else {
					log.Fatal(errors.New("unexpected EOF"))
				}
			}

			name := scanner.Text()
			citizens = append(citizens, NewCitizen(name))
		}

		// check if there is a single winner in terms of unique letters
		By(uniqueLettersDesc).Sort(citizens)
		if citizens[0].NumUniqueLetters > citizens[1].NumUniqueLetters {
			fmt.Printf("Case #%d: %s\n", i+1, citizens[0].Name)
			continue
		}

		// else pick out the tied candidates
		tied := []Citizen{}
		for _, citizen := range citizens {
			if citizen.NumUniqueLetters == citizens[0].NumUniqueLetters {
				tied = append(tied, citizen)
			} else {
				break
			}
		}

		// sort the tied candidates in alphabetical order
		By(alphabetical).Sort(tied)
		fmt.Printf("Case #%d: %s\n", i+1, tied[0].Name)
	}
}
