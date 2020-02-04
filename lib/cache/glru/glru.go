package glru

import (
	"container/list"
	"errors"
	"sync"
)

// GLRU ...
// LRU cache
type GLRU struct {
	size      int
	evictList *list.List
	items     map[interface{}]*item
	lock      sync.Mutex
}

type item struct {
	ele   *list.Element
	value interface{}
}

// NEWGLRU ...
func NEWGLRU(size int) (*GLRU, error) {
	if size <= 0 {
		return nil, errors.New("must provide a positive size")
	}

	lru := &GLRU{
		size:      size,
		evictList: list.New(),
		items:     make(map[interface{}]*item),
	}

	return lru, nil
}

// Set ...
func (lru *GLRU) Set(key interface{}, value interface{}) (evicted bool) {
	lru.lock.Lock()
	defer lru.lock.Unlock()

	// update value
	if v, ok := lru.items[key]; ok {
		v.value = value
		lru.evictList.MoveToFront(v.ele)
		return false
	}

	// set value
	element := lru.evictList.PushFront(key)
	lru.items[key] = &item{
		value: value,
		ele:   element,
	}

	if lru.evictList.Len() > lru.size {
		tail := lru.evictList.Back()
		delete(lru.items, tail.Value)
		lru.evictList.Remove(tail)
		return true
	}

	return false
}

// Get ...
func (lru *GLRU) Get(key interface{}) (value interface{}, ok bool) {
	lru.lock.Lock()
	defer lru.lock.Unlock()

	if v, ok := lru.items[key]; ok { // update value
		lru.evictList.MoveToFront(v.ele)
		return v.value, true
	}

	return nil, false
}

// Contains ...
func (lru *GLRU) Contains(key interface{}) (exist bool) {
	lru.lock.Lock()
	defer lru.lock.Unlock()

	_, exist = lru.items[key]
	return
}

// Peek ...
func (lru *GLRU) Peek(key interface{}) (value interface{}, exist bool) {
	lru.lock.Lock()
	defer lru.lock.Unlock()

	if v, ok := lru.items[key]; ok {
		return v.value, true
	}
	return nil, false
}

// Purge ...
// 清空 cache
func (lru *GLRU) Purge() {
	lru.lock.Lock()
	defer lru.lock.Unlock()

	lru.items = make(map[interface{}]*item)
	lru.evictList = list.New()
}

// Keys  ...
func (lru *GLRU) Keys() []interface{} {
	lru.lock.Lock()
	defer lru.lock.Unlock()

	keys := make([]interface{}, lru.size)
	i := 0
	for tail := lru.evictList.Back(); tail != nil; tail = tail.Prev() {
		keys[i] = tail.Value
		i++
	}
	return keys
}

// Len  ...
func (lru *GLRU) Len() int {
	lru.lock.Lock()
	defer lru.lock.Unlock()

	return lru.size
}
