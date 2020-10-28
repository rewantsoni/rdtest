package trie

import (
	"fmt"
	"strings"
)

type node struct {
	links       [256]*node
	isEndOfWord bool
	wordCount   int
}

func (n *node) markAsEnd() {
	n.isEndOfWord = true
}

func (n *node) incCount() {
	n.wordCount++
}

func newNode() *node {
	return &node{}
}

type Trie struct {
	root *node
}

func NewTrie() *Trie {
	return &Trie{
		root: newNode(),
	}
}

func (t *Trie) Insert(sentence string) {
	words := strings.Split(sentence, " ")

	for _, word := range words {
		insertWord(t, word)
	}

}

func (t *Trie) Print() {
	getH(t.root, "")
}

func getH(curr *node, res string) {
	if curr == nil {
		return
	}

	if curr.isEndOfWord {
		fmt.Printf("%s: %d\n", res, curr.wordCount)
		return
	}

	for i := 0; i < 256; i++ {
		if curr.links[i] != nil {

			getH(curr.links[i], fmt.Sprintf("%s%c", res, rune(i)))
		}
	}
}

func insertWord(t *Trie, word string) {
	curr := t.root

	for _, c := range word {

		if curr.links[c] == nil {
			curr.links[c] = newNode()
		}

		curr = curr.links[c]
	}

	curr.markAsEnd()
	curr.incCount()
}
