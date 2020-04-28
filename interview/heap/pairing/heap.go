package heap

import "fmt"

type HeapElement struct {
	v        int32
	children []*HeapElement
}

type Heap struct {
	root *HeapElement
}

func New() *Heap {
	return &Heap{}
}

func (h *Heap) Peek() (int32, error) {
	if h.root == nil {
		return 0, fmt.Errorf("empty heap")
	}
	return h.root.v, nil
}

func (h *Heap) Push(v int32) {
	h.root = meld(&HeapElement{v: v}, h.root)
}

func meld(h1, h2 *HeapElement) *HeapElement {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	if h1.v < h2.v {
		h1.children = append([]*HeapElement{h2}, h1.children...)
		return h1
	}
	h2.children = append([]*HeapElement{h1}, h2.children...)
	return h2
}

func (h *Heap) Pop() (int32, error) {
	if h.root == nil {
		return 0, fmt.Errorf("empty heap")
	}
	v := h.root.v
	h.root = merge(h.root.children)
	return v, nil
}

func merge(c []*HeapElement) *HeapElement {
	switch len(c) {
	case 0:
		return nil
	case 1:
		return c[0]
	}
	return meld(meld(c[0], c[1]), merge(c[2:]))
}
