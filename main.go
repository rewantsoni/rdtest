package main

import (
	"fmt"
	"redhadtest/trie"
)

func main() {
	testCases := []string{
		`New to Python or choosing between Python 2 and Python 3? Read Python 2 or Python 3.`,
		`Don't communicate by sharing memory, share memory by communicating.`,
		`Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.`,
	}

	for _, testCase := range testCases {
		test(testCase)
	}
}

func test(s string) {
	t := trie.NewTrie()
	t.Insert(s)
	t.Print()
	fmt.Println("________")
}
