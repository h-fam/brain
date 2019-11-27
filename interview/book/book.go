package main

import "fmt"

func pageCount(n int32, p int32) int32 {
	/*
	 * Write your code here.
	 */
	f := p / 2
	b := int32(0)
	if n%2 == 0 {
		b = (n + 1 - p) / 2
	} else {
		b = (n - p) / 2
	}
	if f < b {
		return f
	}
	return b
}

func main() {
	tests := []struct {
		pages  int32
		target int32
	}{
		{pages: 6, target: 2},
		{pages: 5, target: 4},
		{pages: 30, target: 1},
		{pages: 30, target: 2},
		{pages: 30, target: 29},
		{pages: 30, target: 30},
		{pages: 31, target: 30},
		{pages: 31, target: 29},
	}
	for _, tt := range tests {
		fmt.Println(tt, pageCount(tt.pages, tt.target))
	}
}
