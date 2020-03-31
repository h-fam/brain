package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func shift(a []int32) {
    l := len(a)-1
    v := a[l]
    for i := l; i > 0; i-- {
      a[i] = a[i-1]
    }
    a[0] = v
}

func print(a []int32) {
    s := ""
    for _, v := range a {
        s += fmt.Sprintf("%d ", v)
    }
    fmt.Println(s)
}

// Complete the insertionSort2 function below.
func insertionSort2(n int32, arr []int32) {
    for i:=1;i<len(arr);i++{
        found := false
        j := i-1
        for ;j>=0;j-- {
            if arr[j] <= arr[i] {
                break
            }
            found = true
        }
        if found {
            shift(arr[j+1:i+1])
        }
        print(arr)
      }
    }

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    insertionSort2(n, arr)
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

