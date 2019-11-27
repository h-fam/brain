package main

import (
	"fmt"
	"sort"
)

func sortInt32(a []int32) {
	sort.Slice(a, func(i, j int) bool {
		if a[i] > a[j] {
			return false
		}
		return true
	})
}

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	sortInt32(a)
	sortInt32(b)
	ret := int32(0)
	var factors []int32
	for i := a[0]; i <= b[0]; i++ {
		valid := true
		for _, v := range a {
			if i%v != 0 {
				valid = false
				break
			}
		}
		if valid {
			factors = append(factors, i)
		}
	}
	for _, f := range factors {
		valid := true
		for _, val := range b {
			if val%f != 0 {
				valid = false
				break
			}
		}
		if valid {
			fmt.Println("val:", f)
			ret++
		}
	}
	return ret
}

func main() {
	tests := []struct {
		a []int32
		b []int32
	}{
		{a: []int32{3, 2}, b: []int32{12, 18, 36}},
		{a: []int32{2, 4, 8}, b: []int32{16, 64, 96}},
		{a: []int32{3, 4}, b: []int32{24, 48}},
	}
	for _, tt := range tests {
		fmt.Println(tt, getTotalX(tt.a, tt.b))
	}
}
