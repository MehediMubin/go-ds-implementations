package main

import "container/list"

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list list.List
}

type Entry struct {
	key int 
	value int
}

func NewLRUCache (capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache: make(map[int]*list.Element),
		list: *list.New(),
	}
}

