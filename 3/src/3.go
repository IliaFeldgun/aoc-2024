package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

func parseValues() {
  r, _ := regexp.Compile(`(?s)(^(.*?mul\(\d+,\d+\))(?:(.*?don't\(\))))|((?:do\(\).*?)(mul\(\d+,\d+\))(?:.*?don't\(\))?)`)
  input_path := os.Args[1]
  binary_content, err := os.ReadFile(input_path)
  if err != nil {
    log.Fatal(err)
  }

  sum := 0
  line_match := r.FindAllString(string(binary_content), -1)
  for _, line := range line_match {
    log.Print(line)
    sum += getMulSum(line)
  }
  log.Print(sum)
}

func getMulSum(safe_mul_string string) (int) {
  mulr, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
  line_match := mulr.FindAllString(safe_mul_string, -1)
  sum := 0
  for _, mul := range line_match {
    lfactor, rfactor := getFactors(mul)
    sum += lfactor * rfactor
  }
  return sum
}

func getFactors(mul_directive string) (int, int) {
  r, _ := regexp.Compile("[0-9]+")
  factors := r.FindAllString(mul_directive, -1)
  int_factors := []int{}
  for _, v := range factors {
    int_factor, err := strconv.Atoi(v)
    if err != nil {
      log.Fatalf("%s", err)
    }
    int_factors = append(int_factors, int_factor)
  }
  return int_factors[0], int_factors[1]
}

func main() {
  parseValues()
}
