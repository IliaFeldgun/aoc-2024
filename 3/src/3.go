package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func parseValues() {
  input_path := os.Args[1]
  file, err := os.Open(input_path)
  if err != nil {
    log.Fatal(err)
  }
  r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
  lineScanner := bufio.NewScanner(file)
  lineScanner.Split(bufio.ScanLines)
  sum := 0
  for lineScanner.Scan(){
    line_match := r.FindAllString(lineScanner.Text(), -1)
    for _, mul := range line_match {
      lfactor, rfactor := getFactors(mul)
      sum += lfactor * rfactor
    }
  }
  log.Print(sum)
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
