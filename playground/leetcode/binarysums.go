package leetcode

func Binarysearch(arr []int, target int) int {
	s, e := 0, len(arr)-1

	for s <= e {
		m := s + (e-s)/2
		if arr[m] == target {
			return m
		} else if arr[m] > target {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	return s
}

func RunTest() {}
