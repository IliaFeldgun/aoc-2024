package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func parseValues() ([]int, []int){
  input_path := os.Args[1]
  file, err := os.Open(input_path)
  if err != nil {
    log.Fatal(err)
  }
  r, _ := regexp.Compile("[0-9]+")
  larray := []int{}
  rarray := []int{}
  lineScanner := bufio.NewScanner(file)
  lineScanner.Split(bufio.ScanLines)
  for lineScanner.Scan(){
    line_match := r.FindAllString(lineScanner.Text(), -1)
    left, lerr := strconv.Atoi(line_match[0])
    right, rerr := strconv.Atoi(line_match[1])
    if lerr != nil || rerr != nil {
      log.Fatalf("%s%s", lerr, rerr)
    }
    larray = append(larray, left)
    rarray = append(rarray, right)
  }
  return larray, rarray

}
func main() {
  larray, rarray := parseValues()
  slices.SortFunc(larray, func(a, b int) int {return a - b})
  slices.SortFunc(rarray, func(a, b int) int {return a - b})
  sum := 0
  for i := 0; i < len(larray); i++ {
    var diff int
    if larray[i] < rarray[i]{
      diff = rarray[i] - larray[i]
    } else {
      diff = larray[i] - rarray[i]
    }
    sum += diff
    log.Printf("%d ~ %d = %d", larray[i], rarray[i], diff)
  }
  log.Print("Total distance: ", sum)
  score := 0
  count := 0
  for i := 0; i < len(larray); i++ {
    count = 0
    for j := 0; j < len(rarray); j++ {
      if rarray[j] > larray[i] {
        break
      } else if rarray[j] == larray[i] {
        count++
      }
    }
    log.Print(larray[i], " count: ", count)
    score += count * larray[i]
  }
  log.Print(score)
}
