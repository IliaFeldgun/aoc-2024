package aoc_lib

import (
	"log"
	"os"
)

func readArg(index int) string {
	return os.Args[index]
}

func getInput(input_path string) []byte {
	binary_content, err := os.ReadFile(input_path)
	if err != nil {
		log.Fatal(err)
	}
	return binary_content
}

func main() {
	print(string(getInput(readArg(1))))
}
