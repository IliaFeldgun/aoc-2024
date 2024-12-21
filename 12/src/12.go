package main

import (
	"bytes"
	"log"
	"os"
)

const non_char = '.'

type Location struct {
	row int
	col int
}

type Line struct {
	start Location
	end   Location
}

func parseValues() {
	binary_content := getInput()
	bin_matrix := allocate_matrix(binary_content)
	log.Print("\n", string(bytes.Join(bin_matrix, []byte("\n"))))
	plot_region_starts := map[Location]Location{}

	for row := range bin_matrix {
		for col, v2 := range bin_matrix[row] {
			location := Location{row, col}
			plot_region_starts = getRegionStart(bin_matrix, v2, location, plot_region_starts)
		}
	}
	regions := map[Location][]Location{}
	for plot, start := range plot_region_starts {
		region_start, ok := regions[start]
		if !ok {
			regions[start] = []Location{plot}
		} else {
			regions[start] = append(region_start, plot)
		}
	}
	sum := 0
	for region, plots := range regions {
		fence_count := 0
		for _, plot := range plots {
			fence_count += len(getFences(bin_matrix, plot))
		}
		plot_count := len(plots)
		sum += plot_count * fence_count
		log.Print(string(bin_matrix[region.row][region.col]), " ", region, plot_count, " * ", fence_count, " = ", plot_count*fence_count)

	}
	log.Print(sum)
	// sides := map[Line]int{}
	// for region, plots := range regions {
	// 	for _, plot := range plots {
	// 		fences := getFences(bin_matrix, plot)
	// 		for i, l := range fences {
	// 		}
	// 	}
	// }
}

// func getRowFenceLine(bin_matrix [][]byte, f Location) Line {
// 	getSameValueNeighbors(bin_matrix [][]byte, location Location)
//
// }

func applyDiff(l Location, diff Location) Location {
	return Location{l.row + diff.row, l.col + diff.col}
}

func getRegionStart(bin_matrix [][]byte, char byte, location Location, region_starts map[Location]Location) map[Location]Location {
	neighbors := getNeighbors(bin_matrix, location)
	for _, nloc := range neighbors {
		nvalue := bin_matrix[nloc.row][nloc.col]
		if nvalue == char {
			known_start, ok := region_starts[location]
			if !ok {
				region_starts[location] = location
				known_start = region_starts[location]
			}
			n_start, ok := region_starts[nloc]
			if !ok {
				region_starts[nloc] = known_start
				region_starts = getRegionStart(bin_matrix, char, nloc, region_starts)
				n_start = region_starts[nloc]
				known_start = region_starts[location]
			}
			if known_start.row > n_start.row || (known_start.col > n_start.col && known_start.row == n_start.row) {
				region_starts[location] = region_starts[nloc]
			} else if known_start.row < n_start.row || (known_start.col < n_start.col && known_start.row == n_start.row) {
				region_starts[nloc] = region_starts[location]
			}
		}
	}
	_, ok := region_starts[location]
	if !ok {
		region_starts[location] = location
	}
	return region_starts
}

func getFences(bin_matrix [][]byte, location Location) []Location {
	fences := []Location{}
	value := bin_matrix[location.row][location.col]
	neighbors := genNeighbors(location)
	for _, l := range neighbors {
		n_value := getSafeValue(bin_matrix, l)
		if n_value != value {
			fences = append(fences, l)
		}
	}
	return fences
}

func getSameValueNeighbors(bin_matrix [][]byte, location Location) []Location {
	neighbors := getNeighbors(bin_matrix, location)
	good_neighbors := []Location{}
	for _, l := range neighbors {
		if bin_matrix[l.row][l.col] != bin_matrix[location.row][location.col] {
			good_neighbors = append(good_neighbors, Location{l.row, l.col})
		}
	}
	return good_neighbors
}

func getNeighbors(bin_matrix [][]byte, location Location) []Location {
	neighbors := genNeighbors(location)
	good_neighbors := []Location{}
	for _, l := range neighbors {
		value := getSafeValue(bin_matrix, l)
		if value != 0 {
			good_neighbors = append(good_neighbors, Location{l.row, l.col})
		}
	}
	return good_neighbors
}

func genNeighbors(location Location) []Location {
	row := location.row
	col := location.col
	neighbors := []Location{{row - 1, col}, {row, col - 1}, {row + 1, col}, {row, col + 1}}
	return neighbors
}

func getSafeValue(bin_matrix [][]byte, l Location) byte {
	if l.row > -1 && l.col > -1 && l.row < len(bin_matrix) && l.col < len(bin_matrix[0]) {
		return bin_matrix[l.row][l.col]
	}
	return 0
}

func getInput() []byte {
	input_path := os.Args[1]
	binary_content, err := os.ReadFile(input_path)
	if err != nil {
		log.Fatal(err)
	}
	return binary_content
}

func allocate_matrix(bin_content []byte) [][]byte {
	bin_copy := bytes.Clone(bin_content)
	bin_matrix := bytes.Split(bin_copy, []byte("\n"))
	if len(bin_matrix[len(bin_matrix)-1]) == 0 {
		bin_matrix = bin_matrix[:len(bin_matrix)-1]
	}
	return bin_matrix
}

func main() {
	parseValues()
}
