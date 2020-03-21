package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

type bounds struct {
    sX int32
    sY int32
    eX int32
}

func generateSlow(h []int32) int64 {
    max := int64(0)
    for i:=int64(0);i<int64(len(h))-1; i++ {
        for j:=int64(h[i]); j>0; j-- {
            x := i+1
            for x < int64(len(h)){
                if int64(h[x]) < j {
                    break
                }
                x++
            }
            a := (x-i) * j
            if max < a {
                max = a
            }
        }
    }
    return max
}

func generate(h []int32) int64 {
    max := int64(0)
    for i:=int64(0);i<int64(len(h))-1; i++ {
        for j:=int64(h[i]); j>0; j-- {
            x := i+1
            for x < int64(len(h)){
                if int64(h[x]) < j {
                    break
                }
                x++
            }
            a := (x-i) * j
            if max < a {
                max = a
            }
        }
    }
    return max
}

// Complete the largestRectangle function below.
func largestRectangle(h []int32) int64 {
   m := generate(h)
   fmt.Println(m)
   return m
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    hTemp := strings.Split(readLine(reader), " ")

    var h []int32

    for i := 0; i < int(n); i++ {
        hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
        checkError(err)
        hItem := int32(hItemTemp)
        h = append(h, hItem)
    }

    result := largestRectangle(h)

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
