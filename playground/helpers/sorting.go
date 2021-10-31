package helpers

import (
	"fmt"
	"math/rand"
)

func QuickSort(arr []int, start, end int) []int {
	if start < end {
		var p int
		arr, p = partition(arr, start, end)
		// fmt.Printf("arr %v | p %d\n", arr, p)
		QuickSort(arr, start, p-1)
		QuickSort(arr, p+1, end)
	}
	return arr
}

func partition(arr []int, start, end int) ([]int, int) {
	pivot := arr[end]
	idx := start
	for i := start; i < end; i++ {
		if arr[i] < pivot {
			arr[i], arr[idx] = arr[idx], arr[i]
			idx++
		}
	}
	arr[idx], arr[end] = pivot, arr[idx]
	return arr, idx
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := int(len(arr) / 2)
	left := arr[:mid]
	right := arr[mid:]
	// fmt.Printf("left: %v\n", left)
	// fmt.Printf("right: %v\n", right)
	left = MergeSort(left)
	right = MergeSort(right)
	return merge(left, right)
}

func merge(left, right []int) []int {
	nl := len(left)
	nr := len(right)
	result := make([]int, nl+nr)
	l, r, k := 0, 0, 0
	for l < nl && r < nr {
		if left[l] <= right[r] {
			result[k] = left[l]
			l++
		} else {
			result[k] = right[r]
			r++
		}
		k++
	}
	for l < nl {
		result[k] = left[l]
		l++
		k++
	}
	for r < nr {
		result[k] = right[r]
		r++
		k++
	}
	return result
}

func InsertionSort(arr []int) []int {
	//5 4 3 2 1
	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		hole := i
		for hole > 0 && arr[hole-1] > curr {
			arr[hole] = arr[hole-1]
			hole--
		}
		arr[hole] = curr
	}
	return arr
}

func SelectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

func RunTest(size int) {
	arr := rand.Perm(size)
	fmt.Printf("Original arr: %v\n", arr)
	qs := QuickSort(arr, 0, len(arr)-1)
	fmt.Printf("Quick sorted arr: %v\n", qs)
	ms := MergeSort(arr)
	fmt.Printf("merge sorted arr: %v\n", ms)
	ms1 := MergeSort([]int{4, 3, 5, 1, 7})
	fmt.Printf("merge sorted ms1: %v\n", ms1)
	ins := InsertionSort(arr)
	fmt.Printf("Insertion arr: %v\n", ins)
	ss := InsertionSort(arr)
	fmt.Printf("Selection sorted arr: %v\n", ss)
}
