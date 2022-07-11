package main

import (
	"fmt"
)

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}
type Queue struct {
	Head   *Node
	Tail   *Node
	length int
}
type Cache struct {
	Queue Queue
	Hash  Hash
}
type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {

}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()
	for _, word := range []string{"parrot", "avocado", "tree", "potato", "tomato", "cheese", "dog"} {
		cache.Check(word)
		cache.Display()
	}
}
