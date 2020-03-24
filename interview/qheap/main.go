package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "container/heap"
)

type Heap []int64
func (h Heap) Len() int         { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(int64))
}

func (h *Heap) Peek() interface{} {
  return (*h)[0]
}

func (h *Heap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 4096 * 4096)

    tQueries := strings.Split(readLine(reader), " ")

    queries, err := strconv.ParseInt(tQueries[0], 10, 64)
    checkError(err)

    h := &Heap{}
    heap.Init(h)

    for i := int64(0); i < queries; i++ {
        tQuery := strings.Split(readLine(reader), " ")
        qType, err := strconv.ParseInt(tQuery[0], 10, 64)
        checkError(err)
        switch qType {
        case 1:
          v, err := strconv.ParseInt(tQuery[1], 10, 64)
          checkError(err)
          heap.Push(h, v)
        case 2:
          v, err := strconv.ParseInt(tQuery[1], 10, 64)
          checkError(err)
          for i:=0;i<len(*h);i++{
             if v == (*h)[i] {
               *h = append((*h)[:i], (*h)[i+1:]...)
                break
             }

          heap.Init(h)
        case 3:
          fmt.Println(h.Peek())
        }
    }
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
