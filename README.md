# Count Distinct Substrings

> Striver SDE Sheet 188 · Day 27 Trie — insert every suffix; the trie deduplicates shared prefixes structurally, so the node count *is* the answer. Go, stdlib only, tested.

## Problem

How many distinct non-empty substrings does a string have? ("ababa" → 9; naive enumeration materializes O(n²) strings into a set.)

## Approach

Every substring is a prefix of some suffix. Insert all n suffixes into a trie: identical substrings follow identical paths, so each *distinct* substring corresponds to exactly one trie node — count node creations and the deduplication happened structurally, no set, no string allocation. The trie is O(n²) nodes worst case, which is fine at sheet scale; suffix automata or suffix arrays get the same count in O(n log n) when strings grow serious. A brute-force set cross-check over assorted words serves as the property test.

**Complexity:** O(n²) time and space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
