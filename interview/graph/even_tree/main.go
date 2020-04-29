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

type Graph map[int32]Edges

func (g Graph) AddEdge(src, dst int32) {
	n, ok := g[src]
	if !ok {
		n = Edges{}
		g[src] = n
	}
	n[dst] = true
}

func (g Graph) AddNode(n int32) {
	g[n] = Edges{}
}

func (g Graph) String() string {
	s := "Edges:\n"
	var keys []int32
	for v := range g {
		keys = append(keys, v)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	for _, k := range keys {
		var e []string
		for k := range g[k] {
			e = append(e, fmt.Sprintf("%d", k))
		}
		s += fmt.Sprintf("%d: [%s]\n", k, strings.Join(e, ","))
	}
	return s
}

type node struct {
	id       int32
	parent   int32
	children int32
}

func (g Graph) Walk() int32 {
	visited := map[int32]node{}
	g.walk(1, 0, visited)
	cuts := int32(0)
	for _, v := range visited {
		if v.parent == 0 {
			continue
		}
		if v.children%2 == 0 {
			cuts++
		}
	}
	return cuts
}

func (g Graph) walk(n, p int32, visited map[int32]node) int32 {
	curr, ok := visited[n]
	if !ok {
		curr = node{
			id: n,
		}
		visited[n] = curr
	}
	c := int32(1)
	for k := range g[n] {
		if _, ok := visited[k]; ok {
			continue
		}
		c += g.walk(k, n, visited)
	}
	curr.children = c
	curr.parent = p
	visited[n] = curr
	return c
}

type Edges map[int32]bool

// Complete the evenForest function below.
func evenForest(tNodes int32, tEdges int32, tFrom []int32, tTo []int32) int32 {
	g := Graph{}
	for i := int32(1); i <= tNodes; i++ {
		g.AddNode(i)
	}
	for i := range tFrom {
		g.AddEdge(tFrom[i], tTo[i])
		g.AddEdge(tTo[i], tFrom[i])
	}
	fmt.Println(g)
	return g.Walk()
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

	res := evenForest(int32(tNodes), int32(tEdges), tFrom, tTo)

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
