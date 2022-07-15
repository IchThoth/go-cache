package main

import (
	"fmt"
)

const SIZE int = 6

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
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	head.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	c.Add(node)
	c.Hash[str] = node
}
func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove:%s/n", n.Val)

	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	c.Queue.length = -1
	delete(c.Hash, n.Val)
	return n
}
func (c *Cache) Add(n *Node) {
	fmt.Printf("add:%s/n", n.Val)
	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.length++
	if c.Queue.length < SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}
func (c *Cache) Display() {
	c.Queue.Display()
}
func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d-[", q.length)
	for i := 0; i < q.length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()
	for _, word := range []string{"parrot", "avocado", "tree", "potato", "tomato", "cheese", "dog"} {
		cache.Check(word)
		cache.Display()
	}
}
