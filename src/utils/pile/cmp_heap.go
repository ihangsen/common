package pile

import (
	"cmp"
)

type CmpHeap[T cmp.Ordered] []T

func (h CmpHeap[T]) Len() int {
	return len(h)
}

func (h CmpHeap[T]) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h CmpHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *CmpHeap[T]) Push(x T) {
	*h = append(*h, x)
}

func (h *CmpHeap[T]) Pop() T {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
