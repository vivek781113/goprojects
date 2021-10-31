package leetcode

import (
	"fmt"
	"math/rand"
	"sort"
)

func TwoSum(nums []int, target int) []int {
	idxMap := make(map[int]int)
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		compliment := target - curr
		// fmt.Println("iteration: %d | c: %d | tc: %d\n", i, curr, compliment)
		if _, ok := idxMap[compliment]; ok {
			result[0] = i
			result[1] = idxMap[compliment]
			break
		} else {
			idxMap[curr] = i
		}
	}
	// fmt.Println(idxMap)
	return result
}

//merge without extra space
func merge(nums1, nums2 []int) ([]int, []int) {
	i := len(nums1) - 1
	j := 0
	for i >= 0 && j < len(nums2) {
		if nums1[i] > nums2[j] {
			nums1[i], nums2[j] = nums2[j], nums1[i]
		}
		i--
		j++
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	return nums1, nums2
}

func subArrSum(arr []int, k int) int {
	start, currSum, maxSum := 0, 0, 0
	for end := 0; end < len(arr); end++ {
		currSum += arr[end]
		if end-start >= k-1 {
			// fmt.Printf("currSum %d | end %d\n", currSum, end)
			if currSum > maxSum {
				maxSum = currSum
			}
			currSum -= arr[start]
			start += 1

		}
	}
	return maxSum
}

func RunTest1() {

	for i := 0; i < 3; i++ {
		arr := rand.Perm(5)
		fmt.Printf("arr: %v\n", arr)
		fmt.Printf("maxSum: %d\n", subArrSum(arr, 3))
	}

	nums1 := []int{17, 18, 21, 47, 102}
	nums2 := []int{15, 16, 19, 20, 23}

	fmt.Printf("Original nums1: %v\n", nums1)
	fmt.Printf("Original nums2: %v\n", nums2)

	s1, s2 := merge(nums1, nums2)
	fmt.Printf("Sorted nums1: %v\n", s1)
	fmt.Printf("Sorted nums2: %v\n", s2)
	fmt.Printf("############################\n")
}
