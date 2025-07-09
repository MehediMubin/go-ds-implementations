package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list 		 *list.List
}

type Entry struct {
	key int 
	value int
}

func InitLRUCache (capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache: make(map[int]*list.Element),
		list: list.New(),
	}
}

func (lru *LRUCache) Get (key int) {

} 

func (lru *LRUCache) Put (key, value int) {

}

func (lru *LRUCache) Display() {
	
}

func main() {
	var capacity int
	fmt.Println("Enter the capacity of the LRU cache:")
	fmt.Scan(&capacity)

	cache := InitLRUCache(capacity)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1: Get value by key")
		fmt.Println("2: Put key-value pair")
		fmt.Println("0: Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var key int 
			fmt.Print("Enter the key to access: ")
			fmt.Scan(&key)

			cache.Get(key)
			cache.Display()

		case 2:
			var key, value int
			fmt.Print("Enter the key: ")
			fmt.Scan(&key)
			fmt.Print("Enter the value: ")	
			fmt.Scan(&value)

			cache.Put(key, value)
			cache.Display()

		case 0:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Plese try again")
		}
	}
}