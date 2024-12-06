package main

import (
	"bytes"
	"log"
	"os"
)

func getInput() []byte {
	input_path := os.Args[1]
	binary_content, err := os.ReadFile(input_path)
	if err != nil {
		log.Fatal(err)
	}
	return binary_content
}

func parseValues() {
	binary_content := getInput()
	bin_matrix := bytes.Split(binary_content, []byte("\n"))
	guard := "^"
	bin_guard := []byte(guard)
	init_i, init_j := -1, -1
	for i := range bin_matrix {
		for j := range bin_matrix[i] {
			if bin_matrix[i][j] == bin_guard[0] {
				init_i = i
				init_j = j
			}
		}
	}
	if init_i < 0 || init_j < 0 {
		log.Fatal("Guard not found")
	}
	i, j := init_i, init_j
	visited := "X"
	bin_visited := []byte(visited)
	obstacle := "#"
	bin_obstacle := []byte(obstacle)
	direct_i := 0
	// directions := []string{"^", ">", "v", "<"}
	for i > 0 && j > 0 && i < len(bin_matrix) && j < len(bin_matrix[0]) {

		if bin_matrix[i][j] != bin_obstacle[0] {
			bin_matrix[i][j] = bin_visited[0]
			log.Print("Visited ", i, j)
		} else {
			log.Print("Found obstacle ", i, j)
			switch direct_i {
			case 0:
				i++
			case 1:
				j--
			case 2:
				i--
			case 3:
				j++
			}
			direct_i++
			if direct_i == 4 {
				direct_i = 0
			}
			log.Print("Switched direction ", direct_i)
		}
		switch direct_i {
		case 0:
			i--
		case 1:
			j++
		case 2:
			i++
		case 3:
			j--
		}
		log.Print("\n", string(bytes.Join(bin_matrix, []byte("\n"))))
	}
	count := 0
	for row := range bin_matrix {
		for col := range bin_matrix[row] {
			if bin_matrix[row][col] == bin_visited[0] {
				count++
			}
		}
	}
	log.Print(count)
}

func main() {
	parseValues()
}
