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

type Children []pos

func (c Children) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Children) Less(i, j int) bool {
	return c[i].Less(c[j])
}

func (c Children) Len() int {
	return len(c)
}

func (c *Children) Get() Children {
	if c == nil {
		return nil
	}
	return *c
}
func (c *Children) Add(p pos) {
	e := *c
	if len(e) == 0 {
		*c = Children{p}
	}
	for i, v := range e {
		if v == p {
			return
		}
		if p.Less(v) {
			*c = append(e[:i], append(Children{p}, e[i:]...)...)
			return
		}
	}
	*c = append(e, p)
}

func (c *Children) Pop() pos {
	v := (*c)[0]
	*c = (*c)[1:]
	return v
}

type Graph struct {
	nodes map[pos]bool
	edges map[pos]*Children
}

func (g Graph) String() string {
	s := "Edges:\n"
	var sPos Children
	for v := range g.nodes {
		sPos = append(sPos, v)
	}
	sort.Sort(sPos)
	for _, k := range sPos {
		var a []string
		for _, v := range g.edges[k].Get() {
			a = append(a, fmt.Sprintf("%+v", v))
		}
		s += fmt.Sprintf("%+v: [%s]\n", k, strings.Join(a, ","))
	}
	return s
}

func (g *Graph) Add(src, dst pos) {
	if _, ok := g.edges[src]; !ok {
		g.edges[src] = &Children{}
	}
	g.edges[src].Add(dst)
}

type node struct {
	visited  bool
	children map[pos]bool
}

func (g *Graph) find(p pos) Children {
	unvisited := &Children{p}
	var children Children
	visited := map[pos]bool{}
	for unvisited.Len() > 0 {
		c := unvisited.Pop()
		children = append(children, c)
		if visited[c] {
			continue
		}
		visited[c] = true
		for _, v := range g.edges[c].Get() {
			if visited[v] {
				continue
			}
			unvisited.Add(v)
		}
	}
	return children
}

func (g *Graph) Max() int32 {
	var unvisited []pos
	for k := range g.nodes {
		unvisited = append(unvisited, k)
	}
	max := 1
	for _, v := range unvisited {
		connected := g.find(v)
		if len(connected) > max {
			max = len(connected)
		}
	}
	return int32(max)
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
	{-1, 0},
	{1, 0},
	{-1, 1},
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
		edges: map[pos]*Children{},
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
