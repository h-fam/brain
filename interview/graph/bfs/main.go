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
	id   int32
	cost int
}

func (g Graph) paths(r int32) []int32 {
	visited := map[int32]node{}
	q := []node{{
		cost: 0,
		id:   r,
	}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if _, ok := visited[curr.id]; ok {
			continue
		}
		for k := range g[curr.id] {
			if _, ok := visited[k]; ok {
				continue
			}
			q = append(q, node{cost: curr.cost + 1, id: k})
		}
		visited[curr.id] = node{cost: curr.cost}
	}
	fmt.Println(visited)
	paths := make([]int32, len(g))
	for idx := range paths {
		i := int32(idx + 1)
		p, ok := visited[i]
		if !ok {
			paths[idx] = -1
		} else {
			paths[idx] = int32(p.cost) * 6
		}
	}
	return append(paths[:r-1], paths[r:]...)
}

type Edges map[int32]bool

// Complete the bfs function below.
func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	g := Graph{}
	for i := int32(1); i <= n; i++ {
		g.AddNode(i)
	}
	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
		g.AddEdge(edge[1], edge[0])
	}
	fmt.Println(g)
	return g.paths(s)
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
		nm := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nm[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(nm[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		var edges [][]int32
		for i := 0; i < int(m); i++ {
			edgesRowTemp := strings.Split(readLine(reader), " ")

			var edgesRow []int32
			for _, edgesRowItem := range edgesRowTemp {
				edgesItemTemp, err := strconv.ParseInt(edgesRowItem, 10, 64)
				checkError(err)
				edgesItem := int32(edgesItemTemp)
				edgesRow = append(edgesRow, edgesItem)
			}

			if len(edgesRow) != int(2) {
				panic("Bad input")
			}

			edges = append(edges, edgesRow)
		}

		sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		s := int32(sTemp)

		result := bfs(n, m, edges, s)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
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
