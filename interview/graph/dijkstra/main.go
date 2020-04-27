package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Graph struct {
	edges map[int32]Edges
	mu    sync.Mutex
}

func (g *Graph) AddEdge(src, dst, cost int32) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	n, ok := g.edges[src]
	if !ok {
		return fmt.Errorf("node %q not found", src)
	}
	curr, ok := n[dst]
	if !ok || cost < curr {
		n[dst] = cost
	}
	return nil
}

func (g *Graph) AddNode(n int32) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.edges[n] = Edges{}
}

func (g *Graph) String() string {
	g.mu.Lock()
	defer g.mu.Unlock()
	s := "Edges:\n"
	var keys []int32
	for v := range g.edges {
		keys = append(keys, v)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	totalEdges := 0
	for _, k := range keys {
		var e []string
		edges := g.edges[k]
		var ekeys []int32
		for v := range edges {
			ekeys = append(ekeys, v)
		}
		totalEdges += len(edges)
		sort.Slice(ekeys, func(i, j int) bool { return ekeys[i] < ekeys[j] })
		for _, ek := range ekeys {
			e = append(e, fmt.Sprintf("%d:%d", ek, edges[ek]))
		}
		s += fmt.Sprintf("%d: [%s]\n", k, strings.Join(e, ","))
	}
	fmt.Printf("Nodes:%d Edges:%d\n", len(keys), totalEdges)
	return s
}

type node struct {
	id   int32
	cost int32
}

type priorityQueue []node

func (p *priorityQueue) Add(n node) {
	q := *p
	l := len(q)
	if l == 0 {
		*p = append(q, n)
		return
	}
	curr := 0
	if q[curr].cost > n.cost {
		*p = append([]node{n}, q...)
		return
	}
	for curr < l {
		if q[curr].cost > n.cost {
			*p = append(q[:curr], append([]node{n}, q[curr:]...)...)
			return
		}
		curr++
	}
	*p = append(q, n)
}

func (p *priorityQueue) Pop() *node {
	if len(*p) == 0 {
		return nil
	}
	v := (*p)[0]
	*p = (*p)[1:]
	return &v
}

func (g *Graph) paths(r int32) []int32 {
	visited := map[int32]node{}
	q := priorityQueue{{
		cost: 0,
		id:   r,
	}}
	for {
		curr := q.Pop()
		if curr == nil {
			break
		}
		c, ok := visited[curr.id]
		if ok && c.cost < curr.cost {
			continue
		}
		visited[curr.id] = node{cost: curr.cost}
		for k, v := range g.edges[curr.id] {
			d, ok := visited[k]
			c := curr.cost + v
			if !ok || d.cost > c {
				q.Add(node{cost: c, id: k})
			}
		}
	}
	paths := make([]int32, len(g.edges))
	for idx := range paths {
		i := int32(idx + 1)
		p, ok := visited[i]
		if !ok {
			paths[idx] = -1
		} else {
			paths[idx] = p.cost
		}
	}
	return append(paths[:r-1], paths[r:]...)
}

type Edges map[int32]int32

// Complete the shortestReach function below.
func shortestReach(n int32, edges [][]int32, s int32) []int32 {
	g := &Graph{edges: map[int32]Edges{}}
	for i := int32(1); i <= n; i++ {
		g.AddNode(i)
	}
	for _, e := range edges {
		if err := g.AddEdge(e[0], e[1], e[2]); err != nil {
			return nil
		}
		if err := g.AddEdge(e[1], e[0], e[2]); err != nil {
			return nil
		}
	}
	return g.paths(s)
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

			if len(edgesRow) != int(3) {
				panic("Bad input")
			}

			edges = append(edges, edgesRow)
		}

		sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		s := int32(sTemp)

		result := shortestReach(n, edges, s)

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
