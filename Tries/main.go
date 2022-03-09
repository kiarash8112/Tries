package main

import "fmt"

type Node struct {
	value       rune
	children    [26]*Node
	isEndOfWord bool
}
type none struct {
	a string
	b string
	c int
}

func (n *Node) Insert(word string) {
	node := Node{}
	for _, i := range word {
		index := i - rune('a')
		if n.children[index] != nil {
			node = Node{value: i}
			n.children[index] = &node
		}

	}
	node.isEndOfWord = true
}

var b int
var a int

func main() {
	trie := Node{value: 0, isEndOfWord: false}
	trie.Insert("salam")
	fmt.Println(trie.children)
	a = 12

}
