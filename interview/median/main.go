package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "container/heap"
    "math"
)

type MaxHeap []int32
func (h MaxHeap) Len() int         { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int32))
}

func (h *MaxHeap) Peek() int32 {
  return (*h)[0]
}

func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func (h MaxHeap) Size() int {
    return len(h)
}

type MinHeap []int32
func (h MinHeap) Len() int         { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int32))
}

func (h *MinHeap) Peek() int32 {
  return (*h)[0]
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func (h MinHeap) Size() int {
    return len(h)
}

/*
 * Complete the runningMedian function below.
 */
func runningMedian(a []int32) []float64 {
    /*
     * Write your code here.
     */
	largeSide := &MinHeap{}
	smallSide := &MaxHeap{}
	median := float64(a[0])
	medians := []float64{median}
	heap.Push(smallSide, a[0])
    for _, v := range a[1:] {
		if smallSide.Size() > largeSide.Size() {
			if v > smallSide.Peek() {
				heap.Push(largeSide, v)
			} else {
				heap.Push(largeSide, heap.Pop(smallSide))
				heap.Push(smallSide, v)
			}
		} else {
			if v < largeSide.Peek() {
				heap.Push(smallSide, v)
			} else {
				heap.Push(smallSide, heap.Pop(largeSide))
				heap.Push(largeSide, v)
			}
		}
		b := smallSide.Size() - largeSide.Size()
		median := float64(0)
        switch b {
        case 0:
            median = math.Round(((float64(smallSide.Peek() + largeSide.Peek())) / 2) * 10) / 10
        case 1:
            median = float64(smallSide.Peek())
        case -1:
			median = float64(largeSide.Peek())
		default:
			panic("can't be here")
		}
		fmt.Println(median, v, smallSide.Peek(), largeSide.Peek())
        medians = append(medians, median)
    }
    return medians
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024*1024)

    aCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    var a []int32

    for aItr := 0; aItr < int(aCount); aItr++ {
        aItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        aItem := int32(aItemTemp)
        a = append(a, aItem)
    }

    result := runningMedian(a)

    for resultItr, resultItem := range result {
        fmt.Fprintf(writer, "%.1f", resultItem)

        if resultItr != len(result)-1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
