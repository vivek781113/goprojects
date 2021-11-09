package leetcode

import "container/heap"

type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := IntMaxHeap{}
	heap.Init(&h)
	for _, num := range nums {
		heap.Push(&h, num)
	}
	for k > 1 {
		heap.Pop(&h)
		k--
	}
	return heap.Pop(&h).(int)
}
