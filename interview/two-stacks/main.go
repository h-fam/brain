package main
import (
    "fmt"
    "strings"
    "strconv"
    "bufio"
    "io"
    "os"
)

type Stack struct {
    v []int32
}

func (s *Stack) Peek() (int32, error) {
    if len(s.v) == 0 {
        return 0, fmt.Errorf("empty stack")
    }
    return s.v[0], nil
}
func (s *Stack) Push(v int32) {
  s.v = append([]int32{v}, s.v...)
}

func (s *Stack) Pop() (int32, error) {
    if len(s.v) == 0 {
        return 0, fmt.Errorf("empty stack")
    }
    v := s.v[0]
    s.v = s.v[1:]
    return v,nil
}

type ReallyBadQueue struct {
    eStack *Stack
    dStack *Stack
}

func (r *ReallyBadQueue) Enqueue(v int32) {
    r.eStack.Push(v)
}

func (r ReallyBadQueue) Dequeue() (int32, error) {
    if len(r.dStack.v) == 0 {
        r.fill()
    }
    return r.dStack.Pop()
}

func (r ReallyBadQueue) fill() {
    for {
        t, err := r.eStack.Pop()
        if err != nil {
            break
        }
        r.dStack.Push(t)
    }

}
func (r ReallyBadQueue) Peek() (int32, error) {
    if len(r.dStack.v) == 0 {
      r.fill()
    }
    return r.dStack.Peek()
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
   r := &ReallyBadQueue{
        dStack: &Stack{},
        eStack: &Stack{},
   }

    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    queries := int32(qTemp)

    for i := 0; i < int(queries); i++ {
        nmx := strings.Split(readLine(reader), " ")
        qTemp, err := strconv.ParseInt(nmx[0], 10, 64)
        checkError(err)
        q := int32(qTemp)
        switch q {
        case 1:
            vTemp, err := strconv.ParseInt(nmx[1], 10, 64)
            checkError(err)
            r.Enqueue(int32(vTemp))
        case 2:
          r.Dequeue()
        case 3:
          v, _ := r.Peek()
          fmt.Println(v)
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
