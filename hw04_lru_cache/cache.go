package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       *sync.Mutex
}

func (l *lruCache) Inc() {
	l.capacity++
}

func (l *lruCache) Dec() {
	l.capacity--
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	ci := cacheItem{key: key, value: value}

	if _, ok := l.items[key]; ok { // item in the map
		l.items[key] = &ListItem{Value: ci}
		l.queue.PushFront(ci)
		l.Dec()
		return true
	}

	if l.capacity <= 0 { // we need to delete last element for new one
		e := l.queue.Back()
		ci := e.Value.(cacheItem)
		k := ci.key
		l.queue.Remove(e)
		delete(l.items, k)
		l.Inc()
	}

	l.items[key] = &ListItem{Value: ci}
	l.queue.PushFront(ci)
	l.Dec()
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if val, ok := l.items[key]; ok { // item in the map
		l.queue.PushFront(val.Value.(cacheItem))
		return val.Value.(cacheItem).value, true
	}

	return nil, false
}

func (l *lruCache) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mu:       new(sync.Mutex),
	}
}
