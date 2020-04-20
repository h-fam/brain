package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func xorArr(a []int32) int32 {
	if len(a) == 0 {
		return 0
	}
	xor := a[0]
	for _, v := range a[1:] {
		xor = xor ^ v
	}
	return xor
}

func subs(a []int32, l int) [][]int32 {
	var arrs [][]int32
	s := l
	for i := s; i <= len(a); i++ {
		arrs = append(arrs, a[i-l:i])
	}
	return arrs
}

// Complete the sansaXor function below.
func sansaXor(arr []int32) int32 {
	if len(arr)%2 == 0 {
		return 0
	}
	xor := arr[0]
	for i := 2; i < len(arr); i += 2 {
		xor = xor ^ arr[i]
	}
	return xor
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		arrTemp := strings.Split(readLine(reader), " ")

		var arr []int32

		for i := 0; i < int(n); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := sansaXor(arr)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
