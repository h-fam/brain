package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isPow(n int64) bool {
	return (n != 0) && ((n & (n - 1)) == 0)
}

func pow2(n int64) int64 {
	c := int64(0)
	if isPow(n) {
		return n
	}
	for n != 0 {
		n >>= 1
		c += 1
	}

	return 1 << c
}

// Complete the andProduct function below.
func andProduct(a int64, b int64) int64 {
	fmt.Println(strconv.FormatInt(a, 2), strconv.FormatInt(b, 2))
	invert := pow2(b^a) - 1
	fmt.Println(^invert, strconv.FormatInt(^invert, 2))
	return a & ^invert
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	for nItr := 0; nItr < int(n); nItr++ {
		ab := strings.Split(readLine(reader), " ")

		a, err := strconv.ParseInt(ab[0], 10, 64)
		checkError(err)

		b, err := strconv.ParseInt(ab[1], 10, 64)
		checkError(err)

		result := andProduct(a, b)

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
