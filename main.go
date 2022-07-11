package main

import (
	"fmt"
)

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()
	for _, word := range []string{"parrot", "avocado", "tree", "potato", "tomato", "cheese", "dog"} {
		cache.Check(word)
		cache.Display()
	}
}
