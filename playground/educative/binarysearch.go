package educative

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	asc := arr[l]-arr[r] < 0
	for l <= r {
		mid := l + (r-l)/2
		if asc {
			if target == arr[mid] {
				return mid
			} else if target > arr[mid] {
				l = mid + 1
			} else if target < arr[mid] {
				r = mid - 1
			}
		} else {
			if target == arr[mid] {
				return mid
			} else if target > arr[mid] {
				r = mid - 1
			} else if target < arr[mid] {
				l = mid + 1
			}

		}
	}
	return -1
}

func ceilOfNum(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + r - l/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	// fmt.Printf("l: %d | r: %d\n", l, r)
	return l
}

func nextLetter(arr []byte, target byte) byte {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + r - l/2
		if target == arr[m] {
			if m >= len(arr)-1 {
				return arr[0]
			} else {
				return arr[m+1]
			}
		} else if target > arr[m] {
			l = m + 1
		} else if target < arr[m] {
			r = m - 1
		}
	}
	// fmt.Printf("left %v | right %v\n", l, r)
	return arr[l%len(arr)]
}

func searchRange(nums []int, target int) []int {

	// left, right & firstIndex of target
	l, r, fidx := 0, len(nums)-1, -1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			fidx = m
			// keep shrinking right to get the firstIndex of target
			r = m - 1
		} else if target > nums[m] {
			l = m + 1
		} else if target < nums[m] {
			r = m - 1
		}
	}
	//target does not exist return
	if fidx == -1 {
		return []int{-1, -1}
	}

	//reset the boundary
	l, r = fidx, len(nums)-1
	//var to track the last index
	lidx := 0

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			lidx = m
			//keep shrinking left side to get the lastIndex of target
			l = m + 1
		} else if target > nums[m] {
			l = m + 1
		} else if target < nums[m] {
			r = m - 1
		}
	}

	return []int{fidx, lidx}

}

func searchMinDiffElement(nums []int, target int) int {
	if target < nums[0] {
		return nums[0]
	}
	if target > nums[len(nums)-1] {
		return nums[len(nums)-1]
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + r - l/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		}
	}
	// fmt.Printf("l: %d | r: %d\n", l, r)
	if nums[l]-target > target-nums[r] {
		return nums[r]
	} else {
		return nums[l]
	}
}

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + (r-l)/2
		if arr[m] < arr[m+1] {
			l = m + 1
		} else if arr[m] > arr[m+1] {
			r = m
		}
		if l == r {
			break
		}
	}
	return l
}
func searchBitnoicArr(arr []int, target int) int {

	pidx := peakIndexInMountainArray(arr)

	ar1 := arr[:pidx+1]
	ar2 := arr[pidx+1:]

	idx1 := binarySearch(ar1, target)
	if idx1 != -1 {
		return idx1
	}
	idx2 := binarySearch(ar2, target)
	if idx2 != -1 {
		return idx2 + pidx + 1
	}

	return -1
}

func BSRunTest() {
	tcBinarySearch := map[int][]int{}
	tcBinarySearch[10] = []int{4, 6, 10}
	tcBinarySearch[5] = []int{1, 2, 3, 4, 5, 6, 7}
	tcBinarySearch[11] = []int{11, 6, 4}
	tcBinarySearch[4] = []int{10, 6, 4}
	tcBinarySearch[18] = []int{10, 6, 4}

	for k, v := range tcBinarySearch {
		fmt.Printf("Binary search arr %v | target %d | index %d\n", v, k, binarySearch(v, k))
	}
	tcBinarySearch = map[int][]int{}
	tcBinarySearch[8] = []int{1, 2, 4, 5, 6, 7}
	tcBinarySearch[7] = []int{4, 6, 10}
	tcBinarySearch[3] = []int{1, 2, 4, 5, 6, 7}
	tcBinarySearch[0] = []int{1, 2, 4, 5, 6, 7}
	tcBinarySearch[2] = []int{1, 2, 4, 5, 6, 7}
	for k, v := range tcBinarySearch {
		fmt.Printf("Ceil of nums %v | target %d | index %d\n", v, k, ceilOfNum(v, k))
	}

	tcNextLetter := map[byte][]byte{}
	tcNextLetter['f'] = []byte{'a', 'c', 'f', 'h'}
	tcNextLetter['b'] = []byte{'a', 'c', 'f', 'h'}
	tcNextLetter['m'] = []byte{'a', 'c', 'f', 'h'}
	for k, v := range tcNextLetter {
		r := nextLetter(v, k)
		fmt.Printf("charArray: %v| target %c| result: %c\n", v, k, r)
	}

	tcBinarySearch = map[int][]int{}
	tcBinarySearch[8] = []int{5, 7, 7, 8, 8, 10}
	tcBinarySearch[6] = []int{5, 7, 7, 8, 8, 10}
	tcBinarySearch[0] = []int{}
	for k, v := range tcBinarySearch {
		idxs := searchRange(v, k)
		fmt.Printf("searchRange of nums %v | target %d | f-index %d | l-index %d\n", v, k, idxs[0], idxs[1])
	}

	tcBinarySearch = map[int][]int{}
	tcBinarySearch[7] = []int{4, 6, 10}
	tcBinarySearch[4] = []int{4, 6, 10}
	tcBinarySearch[17] = []int{4, 6, 10}
	tcBinarySearch[12] = []int{1, 3, 8, 10, 15}
	for k, v := range tcBinarySearch {
		fmt.Printf("searchMinDiffElement of nums %v | target %d | element %d\n", v, k, searchMinDiffElement(v, k))
	}

	tcBitnoicArr := [][]int{{1, 3, 8, 12, 4, 2}, {3, 8, 3, 1}, {1, 3, 8, 12}, {10, 9, 8}}

	for _, arr := range tcBitnoicArr {
		fmt.Printf("bitonic arr max el index, arr: %v | index %d\n", arr, peakIndexInMountainArray(arr))
	}

	tcBinarySearch = map[int][]int{}
	tcBinarySearch[4] = []int{1, 3, 8, 4, 3}
	tcBinarySearch[8] = []int{3, 8, 3, 1}
	tcBinarySearch[12] = []int{1, 3, 8, 12}
	tcBinarySearch[10] = []int{10, 9, 8}
	for k, v := range tcBinarySearch {
		fmt.Printf("searchBitnoicArr of nums %v | target %d | idx %d\n", v, k, searchBitnoicArr(v, k))
	}

}
