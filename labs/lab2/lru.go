package lru

import (
	"errors"
)

type Cacher interface {
	Get(interface{}) (interface{}, error)
	Put(interface{}, interface{}) error
}

type lruCache struct {
	size      int
	remaining int
	cache     map[string]string
	queue     []string
}

func NewCache(size int) Cacher {
	return &lruCache{size: size, remaining: size, cache: make(map[string]string), queue: make([]string, 0)}
}

func (lru *lruCache) Get(key interface{}) (interface{}, error) {
	// Your code here....
	k := key.(string) // reassign interface value to string
	_, check := lru.cache[k]
	if check { // check if it's in the cache
		lru.queue = append(lru.queue, k) // add the key value to the end of the queue
		return lru.cache[k], nil         // return value at key in cache
	}

	return "-1", errors.New("not found in cache") // if not in cache, return -1 and say 'not found in cache'
}

func (lru *lruCache) Put(key, val interface{}) error {
	// Your code here....
	k := key.(string) // reassign interface value to string
	v := val.(string)

	_, check := lru.cache[k] // check if it's in the cache

	if !check {
		if lru.remaining == 0 { // someone's gotta go cause the cache is full
			leastUsed := lru.queue[0]    // store the least most recently used item in the queue
			lru.qDel(lru.queue[0])       // delete the item at the front of the queue
			delete(lru.cache, leastUsed) // delete the least most recently used item from the map

			lru.queue = append(lru.queue, k) // add to the queue
			lru.cache[k] = v                 // place the 'value' at the 'key' in the cache
		} else { // there is space
			lru.queue = append(lru.queue, k) // add to the queue
			lru.cache[k] = v                 // place the 'value' at the 'key' in the cache

			lru.remaining-- // decrease the 'remaining' spots in the queue
		}
	}
	/*
		lru.cache[k] = v                 // replace the value at the 'key' in the cache.
		lru.queue = append(lru.queue, k) // add to the queue
	*/
	return nil
}

// Delete element from queue
func (lru *lruCache) qDel(ele string) {
	for i := 0; i < len(lru.queue); i++ {
		if lru.queue[i] == ele {
			oldlen := len(lru.queue)
			copy(lru.queue[i:], lru.queue[i+1:])
			lru.queue = lru.queue[:oldlen-1]
			break
		}
	}
}
