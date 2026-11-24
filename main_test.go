package main

import "testing"

// brute counts distinct substrings with a set, for cross-checking.
func brute(s string) int {
	set := map[string]bool{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			set[s[i:j]] = true
		}
	}
	return len(set)
}

func TestCountDistinctSubstrings(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"ababa", 9},
		{"aaa", 3}, // a, aa, aaa
		{"abc", 6},
		{"", 0},
		{"a", 1},
	}
	for _, c := range cases {
		if got := CountDistinctSubstrings(c.in); got != c.want {
			t.Errorf("CountDistinctSubstrings(%q) = %d, want %d", c.in, got, c.want)
		}
	}
	// Property: agree with the brute-force set on assorted strings.
	for _, s := range []string{"banana", "mississippi", "abcabc", "zzzzzz"} {
		if got, want := CountDistinctSubstrings(s), brute(s); got != want {
			t.Errorf("%q: trie %d vs brute %d", s, got, want)
		}
	}
}
