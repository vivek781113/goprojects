package helpers

import "fmt"

//left child: 2 * pos + 1
//rt child: 2 * pos + 2
//parent: pos - 1 /2
type IntHeap struct {
	Nums []int
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *IntHeap) Swap(idx1, idx2 int) {
	h.Nums[idx1], h.Nums[idx2] = h.Nums[idx2], h.Nums[idx1]
}

func (h *IntHeap) Insert(num int) {
	h.Nums = append(h.Nums, num)
	h.HeapifyUp(len(h.Nums) - 1)
}

func (h *IntHeap) HeapifyUp(idx int) {
	for h.Nums[parent(idx)] > h.Nums[idx] {
		h.Swap(parent(idx), idx)
		idx = parent(idx)
	}
}

func (h *IntHeap) Remove() int {
	if len(h.Nums) == 0 {
		return -1
	}
	removed := h.Nums[0]
	last := len(h.Nums) - 1

	h.Nums[0] = h.Nums[last]

	h.Nums = h.Nums[:last]

	h.HeapifyDown(0)
	return removed
}

func (h *IntHeap) HeapifyDown(idx int) {
	lastIdxToCheck := len(h.Nums) - 1
	l, r := left(idx), right(idx)
	var childToCompare int

	for l <= lastIdxToCheck {
		if l == lastIdxToCheck {
			childToCompare = l
		} else if h.Nums[l] < h.Nums[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.Nums[idx] > h.Nums[childToCompare] {
			h.Swap(idx, childToCompare)
			idx = childToCompare
			l, r = left(idx), right(idx)
		} else {
			return
		}
	}
}

func HeapRunTest() {
	h := &IntHeap{}
	h.Insert(10)
	h.Insert(12)
	h.Insert(13)
	h.Insert(1)
	h.Insert(2)
	h.Insert(0)
	h.Insert(10)

	fmt.Println(h)
	fmt.Println(h.Remove())
	fmt.Println(h.Remove())
	fmt.Println(h.Remove())
	fmt.Println(h.Remove())
}
