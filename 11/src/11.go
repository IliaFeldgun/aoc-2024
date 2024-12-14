package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	r, _ := regexp.Compile(`[0-9]+`)

	stones := string(binary_content)
	for range 25 {
		line_match := r.FindAllString(stones, -1)
		next_stones := ""
		for _, stone := range line_match {
			new_stone := stone
			switch {
			case stone == "0":
				new_stone = "1"
			case len(stone)%2 == 0:
				new_stone = split(stone)
			default:
				int_stone, _ := strconv.Atoi(stone)
				int_stone *= 2024
				new_stone = strconv.Itoa(int_stone)
			}
			next_stones = fmt.Sprint(next_stones, " ", new_stone)
		}
		stones = next_stones
	}
	log.Print(stones)
	line_match := r.FindAllString(stones, -1)
	log.Print(len(line_match))
}

func split(stone string) string {
	split_left := stone[:len(stone)/2]
	split_right := stone[len(stone)/2:]
	int_stone_right, err := strconv.Atoi(split_right)
	if err != nil {
		log.Fatal(err)
	}
	split_right = strconv.Itoa(int_stone_right)

	new_stone := fmt.Sprint(split_left, " ", split_right)
	return new_stone
}

func main() {
	parseValues()
}
