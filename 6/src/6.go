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

var (
	bin_guard    = []byte("^")
	bin_obstacle = []byte("#")
	bin_empty    = []byte(".")
	bin_up       = []byte("!")
	bin_right    = []byte(">")
	bin_down     = []byte("v")
	bin_left     = []byte("<")
	directions   = [][]byte{bin_up, bin_right, bin_down, bin_left}
)

func getInput() []byte {
	input_path := os.Args[1]
	binary_content, err := os.ReadFile(input_path)
	if err != nil {
		log.Fatal(err)
	}
	return binary_content
}

func countSteps(matrix [][]byte) int {
	count := 0
	for row := range matrix {
		for col := range matrix[row] {
			for dir := range directions {
				if matrix[row][col] == directions[dir][0] {
					count++
				}
			}
		}
	}
	return count
}

func indexOf(findChar []byte, matrix [][]byte) (int, int) {
	init_i, init_j := -1, -1
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == findChar[0] {
				init_i = i
				init_j = j
			}
		}
	}
	if init_i < 0 || init_j < 0 {
		log.Fatal(string(findChar), " not found in \n", matrix)
	}
	return init_i, init_j
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

func tryObstacle(matrix [][]byte, row int, col int) bool {
	if matrix[row][col] != bin_empty[0] {
		return false
	}
	matrix[row][col] = bin_obstacle[0]
	visits := countVisits(matrix)
	return visits == -1
}

func countVisits(matrix [][]byte) int {
	init_row, init_col := indexOf(bin_guard, matrix)
	matrix[init_row][init_col] = []byte("!")[0]
	row, col := init_row, init_col
	direct_i := UP
	height := len(matrix)
	width := len(matrix[0])
	for row > -1 && col > -1 && row < height && col < width {
		new_row, new_col := moveGuard(row, col, direct_i)
		direct_opposite := rotateClockwise(rotateClockwise(direct_i))
		if new_row == -1 || new_col == -1 || new_row == height || new_col == width {
			break
		}
		switch matrix[new_row][new_col] {
		case bin_empty[0]:
			row = new_row
			col = new_col
			matrix[row][col] = directions[direct_i][0]
		case bin_obstacle[0]:
			direct_i = rotateClockwise(direct_i)
		case directions[direct_i][0]:
			return -1
		case directions[direct_opposite][0]:
			nn_row, nn_col := moveGuard(new_row, new_col, direct_i)
			if matrix[nn_row][nn_col] == bin_obstacle[0] {
				return -1
			} else {
				row = new_row
				col = new_col
				matrix[row][col] = directions[direct_i][0]
			}
		default:
			row = new_row
			col = new_col
			matrix[row][col] = directions[direct_i][0]
		}
	}
	return countSteps(matrix)
}

func parseValues() {
	binary_content := getInput()
	bin_matrix := allocate_matrix(binary_content)
	// init_row, _ := indexOf(bin_guard, bin_matrix)
	visit_count := countVisits(bin_matrix)
	log.Print("Guard total cells visited ", visit_count)
	loop_count := 0
	for row := 0; row < len(bin_matrix); row++ {
		for col := 0; col < len(bin_matrix); col++ {
			if bin_matrix[row][col] != bin_empty[0] && bin_matrix[row][col] != bin_obstacle[0] {
				matrix_copy := allocate_matrix(binary_content)
				if tryObstacle(matrix_copy, row, col) {
					loop_count++
				}
			}
		}
	}
	log.Print("Found loops: ", loop_count)
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
