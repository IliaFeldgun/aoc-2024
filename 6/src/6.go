package main

import (
	"bytes"
	"log"
	"os"
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
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
	direct_i := UP
	// directions := []string{"^", ">", "v", "<"}
	for i > -1 && j > -1 && i < len(bin_matrix)-1 && j < len(bin_matrix[0]) {
		if bin_matrix[i][j] != bin_obstacle[0] {
			bin_matrix[i][j] = bin_visited[0]
			log.Print("Visited ", i, j)
			i, j = moveGuard(i, j, direct_i)
		} else {
			log.Print("Found obstacle ", i, j)
			i, j = backTrack(i, j, direct_i)
			direct_i = rotateClockwise(direct_i)
			log.Print("Switched direction ", direct_i)
		}
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

func rotateClockwise(direction Direction) Direction {
	direction++
	if direction == 4 {
		direction = 0
	}
	return direction
}

func moveGuard(row int, col int, direction Direction) (int, int) {
	switch direction {
	case UP:
		row--
	case RIGHT:
		col++
	case DOWN:
		row++
	case LEFT:
		col--
	}
	return row, col
}

func backTrack(row int, col int, direction Direction) (int, int) {
	switch direction {
	case UP:
		row, col = moveGuard(row, col, DOWN)
	case DOWN:
		row, col = moveGuard(row, col, UP)
	case RIGHT:
		row, col = moveGuard(row, col, LEFT)
	case LEFT:
		row, col = moveGuard(row, col, RIGHT)
	}
	return row, col
}

func main() {
	parseValues()
}
