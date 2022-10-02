package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func minimumKeypadClickCount(text string) int32 {
	// Write your code here
	m := make(map[rune]int)

	// 1. Frequencies
	for _, v := range text {
		m[v] += 1
	}

	global_counter := 0
	counter := 0
	level := 1
	keypad := make(map[rune]int)
	// 2. Put the minimums
	for global_counter < len(text) {
		max := 0
		var key rune
		for k, v := range m {
			// find the max
			if v > max {
				max = v
				key = k
			}
		}
		// insert it in the keypad
		if _, ok := keypad[key]; !ok {
			keypad[key] = level
			counter += 1
			if counter >= 9 {
				level += 1
				counter = 0
			}
		}

		// delete the max from map
		delete(m, key)

		global_counter += 1
	}

	// count the frequencies
	var res int32 = 0
	for _, v := range text {
		res += int32(keypad[v])
	}
	fmt.Println(keypad)
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	text := readLine(reader)

	result := minimumKeypadClickCount(text)

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
