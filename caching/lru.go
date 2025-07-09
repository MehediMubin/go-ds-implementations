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
	element, found := lru.cache[key]
	if found {
		lru.list.MoveToFront(element)
		fmt.Printf("Accessed key %d: %d\n", key, element.Value.(*Entry).value)
	} else {
		fmt.Printf("Key %d not found in cache\n", key)
	}
} 

func (lru *LRUCache) Put (key, value int) {
	element, found := lru.cache[key]
	if found {
		element.Value.(*Entry).value = value
		lru.list.MoveToFront(element)
		fmt.Printf("Updated key %d with value %d\n", key, value)
	} else {
		if lru.list.Len() >= lru.capacity {
			evict := lru.list.Back()
			if evict != nil {
				evictedEntry := evict.Value.(*Entry)
				delete(lru.cache, evictedEntry.key)
				lru.list.Remove(evict)
				fmt.Printf("Evicted key %d\n", evictedEntry.key)
			}
		}
		entry := &Entry{key, value}
		element = lru.list.PushFront(entry)
		lru.cache[key] = element
		fmt.Printf("Added key %d with value %d\n", key, value)
	}
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