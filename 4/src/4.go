package main

import (
	// "bytes"
	"bytes"
	"log"
	"os"
	"regexp"
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
	// r, _ := regexp.Compile(`(XMAS)|(SAMX)`)
	// r, _ := regexp.Compile(`(XMAS)|(SAMX)|(X(?:.{9,11})M(?:.{9,11})A(?:.{9,11}S))|(S(?:.{9,11})A(?:.{9,11})M(?:.{9,11}X))`)
	rxmas, _ := regexp.Compile(`XMAS`)
	rsamx, _ := regexp.Compile(`SAMX`)
	binary_content := getInput()
	by_lines := bytes.Split(binary_content, []byte("\n"))
	log.Print(string(bytes.Join(by_lines, []byte("\n"))))
	by_columns := make([][]byte, len(by_lines[0]))
	diag_length := len(by_lines) + len(by_lines[0])
	by_diagonal := make([][]byte, diag_length)
	by_other_diagonal := make([][]byte, diag_length)
	for diag := 0; diag < diag_length; diag++ {
		by_diagonal[diag] = make([]byte, len(by_lines[0]))
		by_other_diagonal[diag] = make([]byte, len(by_lines[0]))
	}
	for col := 0; col < len(by_lines[0]); col++ {
		// log.Print("making:", col, " ", len(by_lines), by_columns)
		by_columns[col] = make([]byte, len(by_lines))
		for line := 0; line < len(by_lines); line++ {
			diagonal_line := (len(by_diagonal)-1)/2 - col + line
			other_diagonal_line := line + col
			diagonal_col := line
			other_diagonal_col := line
			if line > col {
				diagonal_col = col
			}
			if other_diagonal_line > len(by_other_diagonal)/2-1 {
				other_diagonal_col = len(by_other_diagonal[0]) - 1 - col
			}
			// log.Print("adding: ", " from: ", line, ",", col, " max: ", len(by_lines[line]),
			// 	" to: ", other_diagonal_line, ",", other_diagonal_col,
			// 	" max: ", len(by_other_diagonal), ",", len(by_diagonal[diagonal_line]))
			if len(by_lines[line]) > 0 {
				by_diagonal[diagonal_line][diagonal_col] = by_lines[line][col]
				by_other_diagonal[other_diagonal_line][other_diagonal_col] = by_lines[line][col]
				by_columns[col][line] = by_lines[line][col]
			}
			// log.Print("added")
		}
	}

	everythin := bytes.Join(append(by_lines,
		append(by_columns,
			append(by_other_diagonal,
				by_diagonal...)...)...),
		[]byte("\n"))
	log.Print("\n", string(everythin))
	xmas_match := rxmas.FindAllString(string(everythin), -1)
	samx_match := rsamx.FindAllString(string(everythin), -1)
	log.Print(len(samx_match) + len(xmas_match))
}

func parseValues2() {
	rxmas, _ := regexp.Compile(`(SMASM)|(MSAMS)|(SSAMM)|(MMASS)`)
	binary_content := getInput()
	by_lines := bytes.Split(binary_content, []byte("\n"))
	// three_by_three := [3][3]byte{}
	flattened_three_by_threes := []byte{}
	for line := 0; line < len(by_lines)-3; line++ {
		log.Print("line: ", line)
		for col := 0; col < len(by_lines[line])-2; col++ {
			log.Print("col: ", col, "len: ", string(by_lines[line]))
			for j := 0; j < 3; j++ {
				for k := 0; k < 3; k++ {
					// three_by_three[j][k] = by_lines[line+j][col+k]
					if j+k != 1 && j+k != 3 {
						flattened_three_by_threes = append(flattened_three_by_threes, by_lines[line+j][col+k])
					}
				}
			}
			flattened_three_by_threes = append(flattened_three_by_threes, []byte("XXXXX")...)
		}
	}
	log.Print(string(flattened_three_by_threes))
	xmas_match := rxmas.FindAllString(string(flattened_three_by_threes), -1)
	log.Print(len(xmas_match))
}

func main() {
	parseValues()
	parseValues2()
}
