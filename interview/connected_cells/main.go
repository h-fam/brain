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

type Graph struct {
	nodes map[pos]bool
	edges map[pos][]pos
}

func (g Graph) String() string {
	s := "Edges:\n"
	var sPos []pos
	for v := range g.nodes {
		sPos = append(sPos, v)
	}
	sort.Slice(sPos, func(i, j int) bool { return sPos[i].Less(sPos[j]) })
	for _, k := range sPos {
		var a []string
		for _, v := range g.edges[k] {
			a = append(a, fmt.Sprintf("%+v", v))
		}
		s += fmt.Sprintf("%+v: [%s]\n", k, strings.Join(a, ","))
	}
	return s
}

func (g *Graph) Add(src, dst pos) {
	e := g.edges[src]
	if len(e) == 0 {
		g.edges[src] = []pos{dst}
	}
	for i, v := range e {
		if v == dst {
			return
		}
		if dst.Less(v) {
			g.edges[src] = append(e[:i], append([]pos{dst}, e[i:]...)...)
			return
		}
	}
	g.edges[src] = append(e, dst)
}

func (g *Graph) Max() int32 {
	return 0
}

type pos struct {
	x, y int
}

func (p pos) Less(v pos) bool {
	if p.x == v.x {
		return p.y < v.y
	}
	return p.x < v.x
}

var adjPairs = []pos{
	{-1, -1},
	{0, -1},
	{1, -1},
	{0, -1},
	{0, 1},
	{1, 1},
	{0, 1},
	{1, 1},
}

// Complete the connectedCell function below.
func connectedCell(matrix [][]int32) int32 {
	valid := map[pos]bool{}
	for row, rows := range matrix {
		for col, v := range rows {
			if v == 1 {
				p := pos{row, col}
				valid[p] = true
			}
		}
	}
	g := Graph{
		edges: map[pos][]pos{},
		nodes: valid,
	}
	for p := range valid {
		for _, adj := range adjPairs {
			d := pos{p.x + adj.x, p.y + adj.y}
			if valid[d] {
				g.Add(p, d)
				g.Add(d, p)
			}
		}
	}
	fmt.Println(g)
	return g.Max()
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

	mTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	m := int32(mTemp)

	var matrix [][]int32
	for i := 0; i < int(n); i++ {
		matrixRowTemp := strings.Split(readLine(reader), " ")

		var matrixRow []int32
		for _, matrixRowItem := range matrixRowTemp {
			matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
			checkError(err)
			matrixItem := int32(matrixItemTemp)
			matrixRow = append(matrixRow, matrixItem)
		}

		if len(matrixRow) != int(int(m)) {
			panic("Bad input")
		}

		matrix = append(matrix, matrixRow)
	}

	result := connectedCell(matrix)

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
