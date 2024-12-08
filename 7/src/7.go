package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
)

func reduce(parameters []int, operators []string) int {
	sum := parameters[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case "*":
			sum *= parameters[i+1]
		case "+":
			sum += parameters[i+1]
		}
	}
	return sum
}

func tryMultiply(parameters []int, operators []string, index int) (int, []string) {
	new_operators := make([]string, len(operators))
	copy(new_operators, operators)
	new_operators[index] = "*"
	sum := reduce(parameters, new_operators)
	return sum, new_operators
}

func equalize(result int, parameters []int) bool {
	possible_operators := []string{"+", "*"}
	operator_count := len(parameters) - 1
	permutations := [][]string{}
	slots := make([]int, operator_count)
	done := false
	for !done {
		permutation := make([]string, operator_count)
		for slot, operator_index := range slots {
			permutation[slot] = possible_operators[operator_index]
		}
		// log.Print(permutation)
		// log.Print(permutations)
		permutations = append(permutations, permutation)
		for slot := 0; slot < len(slots); slot++ {
			if slots[slot] < len(possible_operators)-1 {
				slots[slot]++
				break
			} else {
				slots[slot] = 0
			}
			if slot == len(slots)-1 {
				done = true
			}
		}
	}
	for _, permutation := range permutations {
		trial_result := reduce(parameters, permutation)
		if result == trial_result {
			return true
		}
	}
	return false
}

func parseValues() {
	// r, _ := regexp.Compile(`\d+: (\d+ )+`)
	binary_content := getInput()
	bin_matrix := allocate_matrix(binary_content)
	count := 0
	sum := 0
	for row := 0; row < len(bin_matrix); row++ {
		parsed_row := bytes.Split(bin_matrix[row], []byte(": "))
		result, _ := strconv.Atoi(string(parsed_row[0]))
		parameters := bytes.Split(parsed_row[1], []byte(" "))
		int_parameters := []int{}
		for i := range parameters {
			int_param, _ := strconv.Atoi(string(parameters[i]))
			int_parameters = append(int_parameters, int_param)
		}
		if equalize(result, int_parameters) {
			count++
			sum += result
		}
		log.Print(result, " = ", string(bytes.Join(parameters, []byte("+"))))
	}
	log.Print(count, ": ", sum)
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
