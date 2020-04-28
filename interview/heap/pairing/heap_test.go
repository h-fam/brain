package heap

import (
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	tests := []struct {
		desc    string
		in      []int32
		want    []int32
		wantErr bool
	}{{
		desc:    "Empty",
		wantErr: true,
	}, {
		desc: "Inorder",
		in:   []int32{1, 2, 3, 4, 5, 6},
		want: []int32{1, 2, 3, 4, 5, 6},
	}, {
		desc: "desc",
		in:   []int32{6, 5, 4, 3, 2, 1},
		want: []int32{1, 2, 3, 4, 5, 6},
	}, {
		desc: "random",
		in:   []int32{100, 6, 99, 4, 101, 3, 2, 56, 1, 50, 5},
		want: []int32{1, 2, 3, 4, 5, 6, 50, 56, 99, 100, 101},
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			h := New()
			for _, v := range tt.in {
				h.Push(v)
			}
			var got []int32
			for len(got) < len(tt.want) {
				v, err := h.Pop()
				if err != nil {
					if tt.wantErr {
						return
					}
					t.Fatalf("unexpected error: %v", got)
				}
				got = append(got, v)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("heap order invalid: got %v, want %v", got, tt.want)
			}
		})

	}
}
