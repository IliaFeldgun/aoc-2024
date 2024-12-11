package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const non_char = "."

func parseValues() {
	// non_id := []byte(".")[0]
	binary_content := getInput()
	string_content := string(binary_content)
	log.Print(string_content)
	files := make(map[int]int)
	empties := make(map[int]int)
	slotted_disk := []int{}
	id_count := 0
	for i := 0; i < len(string_content); i += 2 {
		files[id_count], _ = strconv.Atoi(string(string_content[i]))
		empties[id_count], _ = strconv.Atoi(string(string_content[i+1]))

		for range files[id_count] {
			slotted_disk = append(slotted_disk, id_count)
		}
		for range empties[id_count] {
			slotted_disk = append(slotted_disk, -1)
		}
		id_count++
	}
	// printDisk(slotted_disk)
	defrag2(slotted_disk)
	printDisk(slotted_disk)
	log.Print(checksum(slotted_disk))
}

func checksum(slotted_disk []int) int {
	sum := 0
	for i, v := range slotted_disk {
		if v != -1 {
			sum += i * v
		}
	}
	return sum
}

func findFileStart(slotted_disk []int, end_index int, file_id int) (int, int) {
	for slotted_disk[end_index] != file_id && end_index >= 0 {
		end_index--
	}
	if end_index < 0 || slotted_disk[end_index] != file_id {
		return -1, -1
	}
	file_start := end_index
	for i := end_index; i > 0 && slotted_disk[i] == file_id; i-- {
		file_start--
	}
	file_start++
	return file_start, end_index
}

func findEmptyEnd(slotted_disk []int, start_index int) (int, int) {
	for slotted_disk[start_index] != -1 {
		start_index++
	}
	empty_end := start_index
	file_id := slotted_disk[start_index]
	if file_id != -1 {
		return -1, -1
	}
	for i := start_index; i < len(slotted_disk) && slotted_disk[i] == file_id; i++ {
		empty_end++
	}
	empty_end--
	return start_index, empty_end
}

func defrag2(slotted_disk []int) {
	last_id := slotted_disk[len(slotted_disk)-1]
	file_index := len(slotted_disk) - 1
	for id := last_id; id > 0; id-- {
		file_start, file_end := findFileStart(slotted_disk, file_index, id)
		file_index = file_start
		for iempty := 0; iempty < file_start; iempty++ {
			empty_start, empty_end := findEmptyEnd(slotted_disk, iempty)
			iempty = empty_end + 1
			file_length := file_end - file_start
			empty_length := empty_end - empty_start
			if file_length <= empty_length && empty_end < file_start {
				for i := 0; i <= file_length; i++ {
					slotted_disk[empty_start+i] = id
					slotted_disk[file_start+i] = -1
				}
				break
			}
		}
	}
}

func printDisk(slotted_disk []int) {
	disk := ""
	for _, v := range slotted_disk {
		slot := strconv.Itoa(v)
		if v == -1 {
			slot = string('.')
		}
		disk = fmt.Sprint(disk, slot)
	}
	log.Print(disk)
}

func getInput() []byte {
	input_path := os.Args[1]
	binary_content, err := os.ReadFile(input_path)
	if err != nil {
		log.Fatal(err)
	}
	return binary_content
}

func main() {
	parseValues()
}
