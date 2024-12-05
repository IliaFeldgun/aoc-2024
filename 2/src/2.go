package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func parseValues() ([][]int){
  input_path := os.Args[1]
  file, err := os.Open(input_path)
  if err != nil {
    log.Fatal(err)
  }
  r, _ := regexp.Compile("[0-9]+")
  lineScanner := bufio.NewScanner(file)
  lineScanner.Split(bufio.ScanLines)
  matrix := [][]int{}
  for lineScanner.Scan(){
    line := []int{}
    line_match := r.FindAllString(lineScanner.Text(), -1)
    for i:= 0; i < len(line_match); i++ {
      number, err := strconv.Atoi(line_match[i])
      if err != nil {
        log.Fatalf("%s", err)
      }
      line = append(line, number)
    }
    matrix = append(matrix, line)
  }
  return matrix
}
func isSafe(line []int) (bool) {
  increasing := line[1] - line[0] > 0
  for j := 0; j < len(line) - 1; j++ {
    current := line[j]
    next := line[j + 1]
    diff := next - current
    unsafe_line := (diff == 0 || diff > 3 || diff < -3) ||
      (diff < 0 && increasing) || 
      (diff > 0 && !increasing)
    if unsafe_line {
      return false
    }
  }
  return true
}
func makeSafe(line []int) (bool) {
  if isSafe(line) {
    return true
  }
  is_safe := false
  for i := 0; i < len(line); i++ {
    lineCopy := make([]int, len(line))
    copy(lineCopy, line)
    line_without_this := append(lineCopy[:i], lineCopy[i+1:]...)
    is_safe = isSafe(line_without_this)
    if is_safe {
      return true
    }
  }
  return false
}
func main() {
  matrix := parseValues()
  safe_lines := 0 
  for i := 0; i < len(matrix); i++ {
    line := matrix[i]
    is_safe := makeSafe(line)
    if is_safe {
      safe_lines++
    }
  }
  log.Print(safe_lines)
}
