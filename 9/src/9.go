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
	log.Print(slotted_disk)
	printDisk(slotted_disk)
	defrag2(slotted_disk)
	printDisk(slotted_disk)
	log.Print(checksum(slotted_disk))
	// for id, count := range files {
	// 	id_char := string(id)
	// 	disk = fmt.Sprint(disk,
	// 		strings.Repeat(id_char, count), strings.Repeat(string(non_id), empties[id]))
	// 	id_count++
	// }
	// defragged := make(map[int][]int)
	// for id := len(files) - 1; id > 0; id-- {
	// 	empty_count := 0
	// 	for id_empty := 0; id_empty < len(empties); id_empty++ {
	// 		defragged_list, ok := defragged[id]
	// 		if !ok {
	// 			defragged_list := []int{}
	// 		}
	// 		defragged_list = append(defragged_list, id_empty)
	// 	}
	// }

	// for i := len(disk) - 1; i >= 0; i-- {
	// 	file_id := disk[i]
	// 	if file_id != non_id {
	// 		inserted := false
	// 		for j := 0; j < len(disk); j++ {
	// 			if disk[j] == non_id && !inserted {
	// 				disk = disk[:j] + string(file_id) + disk[j+1:]
	// 				disk = disk[:i] + non_char + disk[i+1:]
	// 				inserted = true
	// 			}
	// 		}
	//
	// 	}
	// }
	//
	// checksum := 0
	// for i := 1; i < len(disk); i++ {
	// 	value, _ := strconv.Atoi(string(disk[i]))
	// 	checksum += (i - 1) * value
	// }
	// log.Print(checksum)
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

func defrag2(slotted_disk []int) {
	defragged_index := len(slotted_disk) - 1
	for starti := 0; starti <= defragged_index; starti++ {
		if slotted_disk[starti] == -1 {
			for defragged_index > starti {
				if slotted_disk[defragged_index] != -1 {
					slotted_disk[starti] = slotted_disk[defragged_index]
					slotted_disk[defragged_index] = -1
					break
				}
				defragged_index--
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
