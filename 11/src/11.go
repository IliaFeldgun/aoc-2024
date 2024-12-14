package main

import (
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"os"
	"regexp"
	"strconv"
)

// const MAX_LENGTH = 4294967296
const MAX_LENGTH = 2294967296

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
	r, _ := regexp.Compile(`[0-9]+`)
	line_match := r.FindAllString(string(binary_content), -1)
	// int_stones := [MAX_LENGTH]int{}
	int_stones := []int{}
	for _, v := range line_match {
		number, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		// int_stones[i] = number
		int_stones = append(int_stones, number)
	}
	// int_stones[length] = -1
	// last_row := 0
	// last_col := length - 1
	// for _, v := range int_stones {
	count := 0
	// current_length := len(line_match)
	current_length := len(int_stones)
	for range 50 {
		for i := range current_length {
			updated_stone, optional_second := updateStone(int_stones[i])
			int_stones[i] = updated_stone
			if optional_second != -1 {
				int_stones = append(int_stones, optional_second)
				// int_stones[current_length] = optional_second
				current_length++
				// int_stones[current_length] = -1
			}
		}
	}
	count += current_length
	// for _, stone := range init_stones {
	// 	int_stones[0] = stone
	// 	int_stones[1] = -1
	// 	current_length := 1
	// 	for range 25 {
	// 		for i := range current_length {
	// 			updated_stone, optional_second := updateStone(int_stones[i])
	// 			if optional_second != -1 {
	// 				int_stones[current_length] = optional_second
	// 				int_stones[current_length+1] = -1
	// 				current_length++
	// 			}
	// 			int_stones[i] = updated_stone
	// 		}
	// 	}
	// 	count += length
	// }
	log.Print(count)
}

func updateStone(stone int) (int, int) {
	if stone == 0 {
		return 1, -1
	} else {
		digits := 1
		new_stone := stone
		for new_stone > 9 {
			digits++
			new_stone = new_stone / 10
		}
		if digits%2 == 0 {
			return split(stone, digits)
		} else {
			return stone * 2024, -1
		}
	}
}

func split(stone, digits int) (int, int) {
	left := stone / int(math.Pow10(digits/2))
	right := stone % int(math.Pow10(digits/2))
	return left, right
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	parseValues()
}
