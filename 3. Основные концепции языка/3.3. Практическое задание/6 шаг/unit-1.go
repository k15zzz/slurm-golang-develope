package main

import (
	"fmt"
)

type Node[T comparable] struct {
	Key   int
	Value T
	Next  *Node[T]
}

type HashTable[T comparable] struct {
	Table []*Node[T]
	Size  int
}

func NewHashTable[T comparable](size int) *HashTable[T] {
	return &HashTable[T]{
		Table: make([]*Node[T], size),
		Size:  size,
	}
}

func (h *HashTable[T]) Insert(key int, value T) {
	index := h.getHash(key)
	node := &Node[T]{Key: key, Value: value}
	if h.Table[index] == nil {
		h.Table[index] = node
	} else {
		currentNode := h.Table[index]
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = node
	}
}

func (h *HashTable[T]) Get(key int) *Node[T] {
	index := h.getHash(key)
	currentNode := h.Table[index]
	for currentNode != nil {
		if currentNode.Key == key {
			return currentNode
		}
		currentNode = currentNode.Next
	}
	return nil
}

func (h *HashTable[T]) getHash(key int) int {
	return key % h.Size
}

func main() {
	hashTable := NewHashTable[string](10)
	hashTable.Insert(1, "one")
	hashTable.Insert(11, "eleven")
	hashTable.Insert(21, "twenty one")

	fmt.Println(hashTable.Get(1).Value)  // "one"
	fmt.Println(hashTable.Get(11).Value) // "eleven"
	fmt.Println(hashTable.Get(21).Value) // "twenty one"
}
