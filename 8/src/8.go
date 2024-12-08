package main

import (
	"bytes"
	"log"
	"math"
	"os"
)

const non_char = "."

type Location struct {
	row int
	col int
}

func getDistance(a, b Location) Location {
	return Location{
		int(math.Abs(float64(a.row - b.row))),
		int(math.Abs(float64(a.col - b.col))),
	}
}

func parseValues() {
	// r, _ := regexp.Compile(`\d+: (\d+ )+`)
	binary_content := getInput()
	bin_matrix := allocate_matrix(binary_content)
	height := len(bin_matrix)
	width := len(bin_matrix[0])
	log.Print(string(bytes.Join(bin_matrix, []byte("\n"))))
	unique_chars := make(map[byte]bool)
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			char := bin_matrix[row][col]
			if char != []byte(non_char)[0] {
				unique_chars[char] = true
			}
		}
	}
	antennae_locations := make(map[byte][]Location)
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			char := bin_matrix[row][col]
			locations, ok := antennae_locations[char]
			if ok {
				antennae_locations[char] = append(locations, Location{row, col})
			} else {
				if char != []byte(non_char)[0] {
					antennae_locations[char] = make([]Location, len(unique_chars))
					antennae_locations[char] = append(locations, Location{row, col})
				}
			}

		}
	}
	possible_antinode_locations := []Location{}
	for _, locations := range antennae_locations {
		for i := range locations {
			for j := range locations {
				if i != j {
					distance := getDistance(locations[i], locations[j])
					i_row := locations[i].row
					j_row := locations[j].row
					i_col := locations[i].col
					j_col := locations[j].col
					loc_a_row := i_row
					loc_a_col := i_col
					for loc_a_row < height && loc_a_col < width && loc_a_row > -1 && loc_a_col > -1 {
						if i_row > j_row {
							loc_a_row -= distance.row
						} else {
							loc_a_row += distance.row
						}
						if i_col > j_col {
							loc_a_col -= distance.col
						} else {
							loc_a_col += distance.col
						}
						possible_antinode_locations = append(possible_antinode_locations, Location{loc_a_row, loc_a_col})
					}
				}
			}
		}
	}
	locations := possible_antinode_locations
	sure_antinode_locations := []Location{}
	for i := range locations {
		for j := range locations {
			if i != j {
				if (locations[i].row == locations[j].row && locations[i].col == locations[j].col) ||
					(locations[j].row < 0 || locations[j].col < 0 || locations[j].row >= height || locations[j].col >= height) {
					locations[j].col = -1
					locations[j].row = -1
				}
			}
		}
	}
	for i := range locations {
		if locations[i].row != -1 {
			sure_antinode_locations = append(sure_antinode_locations, locations[i])
			bin_matrix[locations[i].row][locations[i].col] = []byte("#")[0]
		}
	}

	log.Print("\n", string(bytes.Join(bin_matrix, []byte("\n"))))
	log.Print(sure_antinode_locations)
	log.Print(len(sure_antinode_locations))
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
