package main

import "fmt"

type trie struct {
	children [26]*trie
}

// CountDistinctSubstrings returns the number of distinct non-empty
// substrings of s. Every substring is a prefix of some suffix, so inserting
// all suffixes into a trie creates exactly one node per distinct substring.
func CountDistinctSubstrings(s string) int {
	root := &trie{}
	count := 0
	for start := 0; start < len(s); start++ {
		n := root
		for i := start; i < len(s); i++ {
			c := s[i] - 'a'
			if n.children[c] == nil {
				n.children[c] = &trie{}
				count++ // a node is born exactly once per distinct substring
			}
			n = n.children[c]
		}
	}
	return count
}

func main() {
	fmt.Println(CountDistinctSubstrings("ababa"))
	fmt.Println(CountDistinctSubstrings("aaa"))
}
