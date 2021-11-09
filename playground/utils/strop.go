package utils

import "sort"

//UniqueNSort function take string as input
//sort & remove the duplicates
func UniqueNSort(word string) string {
	charArray := []byte(word)
	freq := map[byte]bool{}
	for _, char := range charArray {
		if _, ok := freq[char]; ok {
			continue
		} else {
			freq[char] = true
		}
	}
	var sl []byte
	for k := range freq {
		sl = append(sl, k)
	}
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
	return string(sl)
}
