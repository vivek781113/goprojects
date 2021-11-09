package leetcode

import (
	"fmt"
	"go_playground/utils"
)

const Alphabet_Size int = 26

type TrieNode struct {
	children   [Alphabet_Size]*TrieNode
	endOfWords bool
	n          int
}

func getNode() *TrieNode {
	node := &TrieNode{}
	node.endOfWords = false
	for i := 0; i < Alphabet_Size; i++ {
		node.children[i] = nil
	}
	return node
}

func insert(root *TrieNode, key string) {
	temp := root

	for i := 0; i < len(key); i++ {
		idx := key[i] - 'a'
		if temp.children[idx] == nil {
			temp.children[idx] = getNode()
		}
		temp = temp.children[idx]
	}
	temp.n++
	temp.endOfWords = true
}

func search(root *TrieNode, key string) bool {
	temp := root
	for i := 0; i < len(key); i++ {
		idx := key[i] - 'a'
		if temp.children[idx] != nil {
			temp = temp.children[idx]
		} else {
			return false
		}
	}
	return temp != nil && temp.endOfWords
}

func count(node *TrieNode, puzzle, first string, hasSeen bool) int {
	result := 0
	if hasSeen {
		result += node.n
	}

	for i, char := range puzzle {
		idx := char - 'a'
		if node.children[idx] != nil {
			result += count(node.children[idx], puzzle, first, hasSeen || (string(puzzle[i]) == first))
		}

	}
	return result
}

func TrieRunTest() {
	root := getNode()
	words := []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"}

	for i := 0; i < len(words); i++ {
		words[i] = utils.UniqueNSort(words[i])
		if len(words[i]) <= 7 {
			insert(root, words[i])
		}
	}
	fmt.Println(words)
	puzzles := []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}

	for _, puz := range puzzles {
		first := string(puz[0])
		// fmt.Println(first)
		fmt.Printf("key: %s | count: %d\n", puz, count(root, puz, first, false))
	}
}
