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

func (g *Graph) paths(r int32) int32 {
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
		if _, ok := visited[curr.id]; ok {
			continue
		}
		visited[curr.id] = node{cost: curr.cost}
		for k, v := range g.edges[curr.id] {
			if _, ok := visited[k]; ok {
				continue
			}
			q.Add(node{cost: v, id: k})
		}
	}
	cost := int32(0)
	for _, v := range visited {
		cost += v.cost
	}
	return cost
}

type Edges map[int32]int32

/*
 * Complete the 'kruskals' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts WEIGHTED_INTEGER_GRAPH g as parameter.
 */

/*
 * For the weighted graph, <name>:
 *
 * 1. The number of nodes is <name>Nodes.
 * 2. The number of edges is <name>Edges.
 * 3. An edge exists between <name>From[i] and <name>To[i]. The weight of the edge is <name>Weight[i].
 *
 */

func kruskals(gNodes int32, gFrom []int32, gTo []int32, gWeight []int32) int32 {
	g := &Graph{edges: map[int32]Edges{}}
	for i := int32(1); i <= gNodes; i++ {
		g.AddNode(i)
	}
	for i := range gFrom {
		if err := g.AddEdge(gFrom[i], gTo[i], gWeight[i]); err != nil {
			return -1
		}
		if err := g.AddEdge(gTo[i], gFrom[i], gWeight[i]); err != nil {
			return -1
		}
	}

	fmt.Println(g)
	return g.paths(1)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	gNodesEdges := strings.Split(readLine(reader), " ")

	gNodes, err := strconv.ParseInt(gNodesEdges[0], 10, 64)
	checkError(err)

	gEdges, err := strconv.ParseInt(gNodesEdges[1], 10, 64)
	checkError(err)

	var gFrom, gTo, gWeight []int32

	for i := 0; i < int(gEdges); i++ {
		edgeFromToWeight := strings.Split(readLine(reader), " ")

		edgeFrom, err := strconv.ParseInt(edgeFromToWeight[0], 10, 64)
		checkError(err)

		edgeTo, err := strconv.ParseInt(edgeFromToWeight[1], 10, 64)
		checkError(err)

		edgeWeight, err := strconv.ParseInt(edgeFromToWeight[2], 10, 64)
		checkError(err)

		gFrom = append(gFrom, int32(edgeFrom))
		gTo = append(gTo, int32(edgeTo))
		gWeight = append(gWeight, int32(edgeWeight))
	}

	res := kruskals(int32(gNodes), gFrom, gTo, gWeight)

	// Write your code here.
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
