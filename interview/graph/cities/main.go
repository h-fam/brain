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

type Children []int32

func (c Children) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Children) Less(i, j int) bool {
	return c[i] <= c[j]
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

func (c *Children) Add(dst int32) {
	e := *c
	if len(e) == 0 {
		*c = Children{dst}
	}
	for i, v := range e {
		if v == dst {
			return
		}
		if dst < v {
			*c = append(e[:i], append(Children{dst}, e[i:]...)...)
			return
		}
	}
	*c = append(e, dst)
}

func (c *Children) Pop() int32 {
	v := (*c)[0]
	*c = (*c)[1:]
	return v
}

type Node struct {
	id       int32
	children *Children
	cost     int32
}

func (n *Node) Add(dst int32) {
	if n.children == nil {
		n.children = &Children{}
	}
	n.children.Add(dst)
}

func (n Node) String() string {
	var a []string
	for _, v := range n.children.Get() {
		a = append(a, fmt.Sprintf("%+v", v))
	}
	return fmt.Sprintf("%+v: [%s]", n.id, strings.Join(a, ","))
}

type Graph struct {
	nodes map[int32]*Node
	edges []*Node
}

func (g Graph) String() string {
	s := "Edges:\n"
	var sPos Children
	for v := range g.nodes {
		sPos = append(sPos, v)
	}
	sort.Sort(sPos)
	for _, k := range sPos {
		s += g.nodes[k].String() + "\n"
	}
	return s
}

func (g *Graph) AddEdge(src, dst int32) {
	n, ok := g.nodes[src]
	if !ok {
		n = &Node{
			id: src,
		}
		g.nodes[src] = n
	}
	n.Add(dst)
}

func (g *Graph) SortEdges() {
	g.edges = nil
	for _, v := range g.nodes {
		g.edges = append(g.edges, v)
	}
	sort.Slice(g.edges, func(i, j int) bool { return g.edges[i].children.Len() > g.edges[j].children.Len() })
}

func (g *Graph) find(n int32) Children {
	unvisited := &Children{n}
	var children Children
	visited := map[int32]bool{}
	for unvisited.Len() > 0 {
		c := unvisited.Pop()
		children = append(children, c)
		if visited[c] {
			continue
		}
		visited[c] = true
		for _, v := range g.nodes[c].children.Get() {
			if visited[v] {
				continue
			}
			unvisited.Add(v)
		}
	}
	return children
}

// Complete the roadsAndLibraries function below.
func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	if c_lib <= c_road {
		return int64(n * c_lib)
	}
	g := &Graph{nodes: map[int32]*Node{}}
	for _, edge := range cities {
		g.AddEdge(edge[0], edge[1])
		g.AddEdge(edge[1], edge[0])
	}
	g.SortEdges()
	fmt.Println(g)
	fmt.Println(g.edges)
	return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nmC_libC_road := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nmC_libC_road[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(nmC_libC_road[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		c_libTemp, err := strconv.ParseInt(nmC_libC_road[2], 10, 64)
		checkError(err)
		c_lib := int32(c_libTemp)

		c_roadTemp, err := strconv.ParseInt(nmC_libC_road[3], 10, 64)
		checkError(err)
		c_road := int32(c_roadTemp)

		var cities [][]int32
		for i := 0; i < int(m); i++ {
			citiesRowTemp := strings.Split(readLine(reader), " ")

			var citiesRow []int32
			for _, citiesRowItem := range citiesRowTemp {
				citiesItemTemp, err := strconv.ParseInt(citiesRowItem, 10, 64)
				checkError(err)
				citiesItem := int32(citiesItemTemp)
				citiesRow = append(citiesRow, citiesItem)
			}

			if len(citiesRow) != 2 {
				panic("Bad input")
			}

			cities = append(cities, citiesRow)
		}

		result := roadsAndLibraries(n, c_lib, c_road, cities)

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
