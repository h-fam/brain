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

/*
 * Complete the cookies function below.
 */

type Heap []int32
func (h Heap) Len() int         { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
    *h = append(*h, x.(int32))
}

func (h *Heap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func cookies(k int32, A []int32) int32 {
    /*
     * Write your code here.
     */
   h := &Heap{}
   *h = Heap(A)
   heap.Init(h)
   ops := int32(0)
   for {
       if len(*h) < 1 {
           return -1
       }
       c1 := heap.Pop(h).(int32)
       if c1 >= k {
           return ops
       }
       if len(*h) < 1 {
           return -1
       }
       c2 := heap.Pop(h).(int32)
       c1 = c1 + c2 * 2
       heap.Push(h, c1)
       ops++
   }
   return ops
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 4096 * 4096)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout,  1024 * 1024)

    nk := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    ATemp := strings.Split(readLine(reader), " ")

    var A []int32

    for AItr := 0; AItr < int(n); AItr++ {
        AItemTemp, err := strconv.ParseInt(ATemp[AItr], 10, 64)
        checkError(err)
        AItem := int32(AItemTemp)
        A = append(A, AItem)
    }

    result := cookies(k, A)

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
