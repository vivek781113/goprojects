package leetcode

import (
	"go_playground/helpers"
	"go_playground/utils"
)

//1178. Number of Valid Words for Each Puzzle
func findNumOfValidWords(words []string, puzzles []string) []int {
	root := helpers.GetNode()
	for _, word := range words {
		word = utils.UniqueNSort(word)
		if len(word) <= 7 {
			helpers.Insert(root, word)
		}
	}

	return []int{}
}

// func

func TrieRunTest() {

}
