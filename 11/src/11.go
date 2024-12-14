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
const MAX_LENGTH = 294967296

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
	int_map := map[int]int{}
	for i, v := range line_match {
		number, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		_, ok := int_map[i]
		if !ok {
			int_map[number] = 1
		} else {
			int_map[number] += 1
		}
	}
	for range 25 {
		to_update := map[int]int{}
		for k, v := range int_map {
			stone := k
			count := v
			updated_stone, optional_second := updateStone(stone)
			_, ok := to_update[updated_stone]
			if !ok {
				to_update[updated_stone] = count
			} else {
				to_update[updated_stone] += count
			}
			if optional_second != -1 {
				_, ok := to_update[optional_second]
				if !ok {
					to_update[optional_second] = count
				} else {
					to_update[optional_second] += count
				}
			}
		}
		int_map = to_update
	}
	count := 0
	for _, v := range int_map {
		count += v
	}
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
