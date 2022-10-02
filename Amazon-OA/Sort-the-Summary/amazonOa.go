package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'groupSort' function below.
 *
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func groupSort(arr []int32) [][]int32 {
	// Write your code here

	// [value, frequency]
	m := make(map[int32]int32)

	// 1. Count the occurrences
	for _, v := range arr {
		m[v] += 1
	}

	fmt.Println(m)
	res := [][]int32{}
	// 2. If freq is equal sort by value
	// O(N^2)
	cond := true
	for cond == true {
		arr := []int32{}
		var max int32 = 0
		for k, v := range m {
			if v > max {
				max = v
				arr = []int32{}
				arr = append(arr, k)
			} else if v == max {
				arr = append(arr, k)
			}
		}
		// sort by value
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
		// insert into res
		for _, v := range arr {
			res = append(res, []int32{v, m[v]})
			// delete from map
			delete(m, v)
		}
		// check if we are done
		if len(m) == 0 {
			cond = false
		}
		fmt.Println(arr)
	}
	fmt.Printf("res = %v\n", res)
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := groupSort(arr)

	for i, rowItem := range result {
		for j, colItem := range rowItem {
			fmt.Fprintf(writer, "%d", colItem)

			if j != len(rowItem)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
