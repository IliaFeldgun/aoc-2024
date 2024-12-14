package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const non_char = "."

type Location struct {
	row int
	col int
}

func parseValues() {
	binary_content := getInput()
	bin_matrix := allocate_matrix(binary_content)
	height := len(bin_matrix)
	width := len(bin_matrix[0])
	log.Print(string(bytes.Join(bin_matrix, []byte("\n"))))
	paths := make(map[Location]map[Location]bool)
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if bin_matrix[row][col] == []byte("0")[0] {
				this_location := Location{row, col}
				this_neighbors := getNextNeighbors(bin_matrix, this_location)
				this_neighbors[this_location] = true
				paths[Location{row, col}] = this_neighbors
			}
		}
	}
	log.Print(paths)
	sum := 0
	for l, v := range paths {
		nine_score := 0

		if bin_matrix[l.row][l.col] == []byte("0")[0] {
			// if bin_matrix[l.row][l.col] == []byte("0")[0] &&
			// 	l.row == 0 || l.col == 0 || l.row == len(bin_matrix)-1 || l.col == len(bin_matrix[0])-1 {
			for i := range v {
				if bin_matrix[i.row][i.col] == []byte("9")[0] {
					nine_score++
				}
			}
		}
		if nine_score > 0 {
			sum += nine_score
			printPath2(v, bin_matrix)
			log.Print(nine_score)
		}
	}
	log.Print(len(paths))
	log.Print(sum)
	// for i := []byte("0")[0]; i < []byte("9")[0]; i++ {
	// 	height_locations
	// 	for j := []byte("0")[0]; j < []byte("9")[0]; j++ {
	//
	// 	}
	// }
}

// func followPath(bin_matrix [][]byte, paths map[Location][]Location, location Location) []Location {
// }

func printPath(path []Location, bin_matrix [][]byte) {
	log.Print(len(path))
	output := ""
	for _, v := range path {
		output = fmt.Sprint(output, " ", string(bin_matrix[v.row][v.col]))
	}

	log.Print(output)
}

func printPath2(path map[Location]bool, bin_matrix [][]byte) {
	output_matrix := allocate_matrix(bytes.Join(bin_matrix, []byte("\n")))
	for row := range output_matrix {
		for col := range output_matrix[row] {
			_, ok := path[Location{row, col}]
			if !ok {
				output_matrix[row][col] = []byte(".")[0]
			}
		}
	}
	log.Print("\n", string(bytes.Join(output_matrix, []byte("\n"))))
}

func getNextNeighbors(bin_matrix [][]byte, location Location) map[Location]bool {
	if bin_matrix[location.row][location.col] == []byte("9")[0] {
		return map[Location]bool{location: true}
	}
	neighbors := getNeighbors(bin_matrix, location)
	next_neighbors := map[Location]bool{}
	for _, l := range neighbors {
		if bin_matrix[location.row][location.col] == bin_matrix[l.row][l.col]-1 {
			next_neighbors[l] = true
			for k, v := range getNextNeighbors(bin_matrix, l) {
				next_neighbors[k] = v
			}
		}
	}
	return next_neighbors
}

func getNeighbors(bin_matrix [][]byte, location Location) []Location {
	row := location.row
	col := location.col
	neighbors := []Location{{row - 1, col}, {row + 1, col}, {row, col - 1}, {row, col + 1}}
	good_neighbors := []Location{}
	for _, l := range neighbors {
		if l.row > -1 && l.col > -1 && l.row < len(bin_matrix) && l.col < len(bin_matrix[0]) {
			good_neighbors = append(good_neighbors, Location{row: l.row, col: l.col})
		}
	}
	return good_neighbors
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
