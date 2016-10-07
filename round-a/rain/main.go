package main

import (
	"bufio"
	"fmt"
	"github.com/yi-jiayu/google-apac-test-2017/codejam"
	"log"
	"os"
)

// Algorithm
// 1. Starting with h = 1
// 2. Loop through all cells
// 3. Once a cell of height h is encountered, run a flood fill starting from that cell
// 3.1. If a boundary cell with height lower than h is found, set flag drain = true
// 3.2. Once the flood fill is finished, if drain = true then mark all cells in the flood fill as visited
// 3.3 Otherwise, increment the height of every cell inside the flood fill by one
// 4. Continue looping through the cells and repeat step 3 if necessary, skipping visited cells
// 5. After looping through all the cells, set h = h + 1 and continue from step 2

type Island struct {
	R     int
	C     int
	Cells [][]int
}

func (i Island) maxHeight() int {
	max := 0

	for _, row := range i.Cells {
		for _, col := range row {
			if col > max {
				max = col
			}
		}
	}

	return max
}

func floodFill(island *Island, h, r, c int) (bool, [][2]int) {
	drain := false
	vis := make(map[[2]int]struct{})
	visited := [][2]int{}

	recursiveFloodFill(island, h, r, c, &drain, &vis)

	for k := range vis {
		visited = append(visited, k)
	}

	return drain, visited
}

func recursiveFloodFill(island *Island, h, r, c int, drain *bool, visited *map[[2]int]struct{}) {
	if r >= island.R || r < 0 || c >= island.C || c < 0 {
		*drain = true
		return
	}

	if _, exists := (*visited)[[2]int{r, c}]; exists {
		return
	}

	if island.Cells[r][c] < h {
		*drain = true
		return
	} else if island.Cells[r][c] > h {
		return
	}

	(*visited)[[2]int{r, c}] = struct{}{}

	recursiveFloodFill(island, h, r+1, c, drain, visited)
	recursiveFloodFill(island, h, r-1, c, drain, visited)
	recursiveFloodFill(island, h, r, c+1, drain, visited)
	recursiveFloodFill(island, h, r, c-1, drain, visited)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Find number of test cases
	T, err := codejam.GetNumTestCases(scanner)
	if err != nil {
		log.Fatal(err)
	}

	// For each test case
	for i := 0; i < T; i++ {
		// Get the size of the island
		RC, err := codejam.ReadLineOfInts(scanner)
		if err != nil {
			log.Fatal(err)
		}

		R, C := RC[0], RC[1]
		island := Island{
			R,
			C,
			[][]int{},
		}

		// Read the heights of the island
		for j := 0; j < R; j++ {
			row, err := codejam.ReadLineOfInts(scanner)
			if err != nil {
				log.Fatal(err)
			}

			island.Cells = append(island.Cells, row)
		}

		increment := 0

		// Start from height 1
		for h := 1; h < island.maxHeight(); h++ {
			// Keep track of which cells we have already visited at this level
			visited := make(map[[2]int]struct{})

			// Loop through all the cells of the island
			for r := 0; r < R; r++ {
				for c := 0; c < C; c++ {
					// Skip if we have already visited it
					if _, exists := visited[[2]int{r, c}]; exists {
						continue
					}

					// Skip if it is not at our current height
					if island.Cells[r][c] != h {
						continue
					}

					// Once we find a cell at our current height, flood fill it to find out if it
					// will fill or drain
					drained, vis := floodFill(&island, h, r, c)

					// For each cell in this block of same-height cells
					for _, coords := range vis {
						r, c := coords[0], coords[1]
						// If the block will fill, increment the height of every cell in the
						// block by one
						if !drained {
							island.Cells[r][c]++
							increment++
						}

						// Record all the cells in this block as visited
						visited[coords] = struct{}{}
					}
				}
			}
		}

		fmt.Printf("Case #%d: %d\n", i+1, increment)
	}
}
