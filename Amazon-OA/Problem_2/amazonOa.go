package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getMaximumGreyness' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING_ARRAY pixels as parameter.
 */

func getMaximumGreyness(pixels []string) int32 {
	// Write your code here
	x := len(pixels)
	y := len(pixels[0])
	max := -1 << 31
	white_cells := 0
	black_cells := 0
	rows := make(map[int][]int)
	cols := make(map[int][]int)

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			white_cells = 0
			black_cells = 0
			// row count
			if _, ok := rows[j]; !ok {
				w_cells := 0
				b_cells := 0
				for i_aux := 0; i_aux < x; i_aux++ {
					val := pixels[i_aux][j]
					if val == byte('1') {
						w_cells += 1
					} else {
						b_cells += 1
					}
				}
				rows[j] = []int{w_cells, b_cells}
			}
			white_cells += rows[j][0]
			black_cells += rows[j][1]

			if _, ok := cols[i]; !ok {
				w_cells := 0
				b_cells := 0
				for j_aux := 0; j_aux < y; j_aux++ {
					val := pixels[i][j_aux]
					if val == byte('1') {
						w_cells += 1
					} else {
						b_cells += 1
					}
				}
				cols[i] = []int{w_cells, b_cells}
			}
			white_cells += cols[i][0]
			black_cells += cols[i][1]

			greyness := white_cells - black_cells
			if greyness > max {
				max = greyness
			}
		}
	}
	return int32(max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	pixelsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var pixels []string

	for i := 0; i < int(pixelsCount); i++ {
		pixelsItem := readLine(reader)
		pixels = append(pixels, pixelsItem)
	}

	result := getMaximumGreyness(pixels)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
