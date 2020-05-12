package main

import (
	"fmt"
	"strings"
)

type node struct {
	next *node
}

type sets struct {
	sets    map[*node]map[*node]bool
	members map[*node]*node
}

func (s sets) String() string {
	out := "sets:\n"
	var elem []string
	for k, v := range s.sets {
		var e []string
		for k := range v {
			e = append(e, fmt.Sprintf("%p", k))
		}
		elem = append(elem, fmt.Sprintf(" %p:[%s]", k, strings.Join(e, ", ")))
	}
	out += strings.Join(elem, "\n")
	out += "\nmembers:\n"
	elem = nil
	for k, dSet := range s.members {
		elem = append(elem, fmt.Sprintf(" %p:%p", k, dSet))
	}
	out += strings.Join(elem, "\n")
	out += "\n"
	return out
}

func (s sets) merge(n1, n2 *node) {
	dKey := s.members[n1]
	sKey := s.members[n2]
	for k := range s.sets[sKey] {
		s.sets[dKey][k] = true
		s.members[k] = dKey
	}
	if dKey != sKey {
		delete(s.sets, sKey)
	}
}

func Set(r *node, set []*node) int {
	s := sets{
		sets:    map[*node]map[*node]bool{},
		members: map[*node]*node{},
	}
	for _, v := range set {
		s.sets[v] = map[*node]bool{v: true}
		s.members[v] = v
	}
	for _, curr := range set {
		dSet := s.members[curr]
		if curr != dSet {
			continue
		}
		next := curr.next
		for next != nil {
			cm := s.members[next]
			if cm == nil {
				break
			}
			s.merge(curr, next)
			next = next.next
		}
	}
	fmt.Println(s)
	return len(s.sets)
}

func main() {
	n10 := &node{}
	n9 := &node{next: n10}
	n8 := &node{next: n9}
	n7 := &node{next: n8}
	n6 := &node{next: n7}
	n5 := &node{next: n6}
	n4 := &node{next: n5}
	n3 := &node{next: n4}
	n2 := &node{next: n3}
	n1 := &node{next: n2}
	set := []*node{n6, n3, n4, n1, n8, n10}
	curr := n1
	i := 0
	for curr != nil {
		fmt.Printf("%d:%p ", i, curr)
		i++
		curr = curr.next
	}
	fmt.Println()
	fmt.Println(Set(n1, set))
}
