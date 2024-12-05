package main

import (
	// "bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
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
	r2, _ := regexp.Compile(`(\d+,).*`)
	r2_match := r2.FindAllString(string(binary_content), -1)
	log.Print(r2_match)
	sum := 0
	sum_corrected := 0
	for _, v := range r2_match {
		pages := strings.Split(v, ",")
		bad_page := false
		for i := 0; i < len(pages); i++ {
			for j := 0; j < len(pages); j++ {
				if pages[j] == "" || pages[i] == "" || bad_page {
					break
				}
				rpages, _ := regexp.Compile(fmt.Sprintf("%s\\|%s", pages[i], pages[j]))
				if rpages.Match(binary_content) && j < i {
					log.Print(rpages.FindAllString(string(binary_content), -1))
					log.Print("Bad: ", pages)
					bad_page = true
				}

			}
		}
		if bad_page {
			sort.Slice(pages, func(i, j int) bool {
				rpages, _ := regexp.Compile(fmt.Sprintf("%s\\|%s", pages[i], pages[j]))
				return rpages.Match(binary_content)
			})
			middle_page, err := strconv.Atoi(pages[(len(pages)-1)/2])
			if err == nil {
				sum_corrected += middle_page
				log.Print("Corrected: ", pages)
			}
		} else {
			middle_page, err := strconv.Atoi(pages[(len(pages)-1)/2])
			if err == nil {
				sum += middle_page
				log.Print("Good: ", pages)
			}
		}
	}
	log.Print("Good sum: ", sum)
	log.Print("Corrected sum: ", sum_corrected)
}

func main() {
	parseValues()
}
