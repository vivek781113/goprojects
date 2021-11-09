package leetcode

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

func count(node *TrieNode, puzzle, first string, start int, met_first bool) int {
	result := 0
	if met_first {
		result += node.n
	}
	for i := start; i < len(puzzle); i++ {
		if node.children[puzzle[i]-'a'] == nil {
			continue
		}
		result += count(node.children[puzzle[i]-'a'], puzzle, first, i+1, (met_first || (string(puzzle[i]) == first)))
	}
	return result
}

func TrieRunTest() {
	words := []string{"a", "as", "abel", "abilty", "act", "acort", "aces"}
	root := getNode()

	for _, word := range words {
		insert(root, word)
	}
	// puzzles := []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}

	// for _, puz := range puzzles {
	// 	first := string(puz[0])
	// 	// fmt.Println(first)
	// 	fmt.Printf("key: %s | exist: %d\n", puz, count(root, puz, first, 0, false))
	// }
	puz := "abslute"
	first := string(puz[0])
	count(root, puz, first, 0, false)
}
