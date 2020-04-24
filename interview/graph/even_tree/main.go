package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the evenForest function below.
func evenForest(t_nodes int32, t_edges int32, t_from []int32, t_to []int32) int32 {

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tNodesEdges := strings.Split(readLine(reader), " ")

	tNodes, err := strconv.ParseInt(tNodesEdges[0], 10, 64)
	checkError(err)

	tEdges, err := strconv.ParseInt(tNodesEdges[1], 10, 64)
	checkError(err)

	var tFrom, tTo []int32

	for i := 0; i < int(tEdges); i++ {
		edgeFromToWeight := strings.Split(readLine(reader), " ")

		edgeFrom, err := strconv.ParseInt(edgeFromToWeight[0], 10, 64)
		checkError(err)

		edgeTo, err := strconv.ParseInt(edgeFromToWeight[1], 10, 64)
		checkError(err)

		tFrom = append(tFrom, int32(edgeFrom))
		tTo = append(tTo, int32(edgeTo))
	}

	res := evenForest(t_nodes, t_edges, t_from, t_to)

	fmt.Fprintf(writer, "%d\n", res)

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
