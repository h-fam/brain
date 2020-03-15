package main
import (
    "fmt"
    "strings"
)

type Stack struct {
    v []int32
}
func (s *Stack) Push(v int32) {
  s.v = append([]int32{v}, s.v...)
}

func (s *Stack) Pop() int32 {
  v := s.v[0]
  s.v = s.v[1:]
  return v
}

type ReallyBadQueue struct {
    s1 *Stack
    s2 *Stack
}

func (r ReallyBadQueue) Enqueue(v int32) {
    r.s1.Push(v)
}

func (r ReallyBadQueue) Dequeue() int32 {
  return r.s1.Pop()
}

func (r ReallyBadQueue) Peek() int32 {
    v := r.s1.Pop()
    r.s1.Push(v)
    return v
}

func (r ReallyBadQueue) String() string {
    s := "{s1:"
    var vals []string
    for _, v := range r.s1.v {
        vals = append(vals, fmt.Sprintf("%d", v))
    }
    s = fmt.Sprintf("%s[%s], s2:", s, strings.Join(vals, ","))
    vals = nil
    for _, v := range r.s2.v {
        vals = append(vals, fmt.Sprintf("%d", v))
    }
    s = fmt.Sprintf("%s[%s]}", s, strings.Join(vals, ","))
    return s
}
func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
 r := ReallyBadQueue{
     s1: &Stack{},
     s2: &Stack{},
 }
 r.Enqueue(1)
 fmt.Println(r, r.Peek())
 r.Enqueue(2)
 fmt.Println(r, r.Peek())
 r.Enqueue(3)
 fmt.Println(r, r.Peek())
 r.Dequeue()
 fmt.Println(r, r.Peek())
}
