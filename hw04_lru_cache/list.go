package hw04lrucache

import (
	"sync"
)

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	mu   *sync.Mutex
	head *ListItem
	tail *ListItem
}

func NewList() List {
	list := new(list)
	list.mu = new(sync.Mutex)
	return list
}

func (l list) Len() int {
	return l.len
}

func (l *list) Inc() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.len++
}

func (l *list) Dec() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.len--
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := &ListItem{Value: v}

	if l.head == nil { // empty dll
		l.head = newListItem
		l.tail = newListItem
	} else {
		newListItem.Next = l.head
		l.head.Prev = newListItem
		l.head = newListItem
	}
	l.Inc()

	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.tail == nil {
		l.tail = &ListItem{Value: v, Next: nil, Prev: nil}
		l.head = l.tail
	} else {
		prevTail := l.tail
		l.tail = &ListItem{Value: v, Next: nil, Prev: prevTail}
		prevTail.Next = l.tail
	}

	l.Inc()

	return l.tail
}

func (l *list) Remove(i *ListItem) {
	defer l.Dec()

	if l.len == 1 {
		l.head = nil
		l.tail = nil
		return
	}

	switch {
	case i.Prev == nil:
		i.Next.Prev = nil
		l.head = i.Next
	case i.Next == nil:
		i.Prev.Next = nil
		l.tail = i.Prev
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil { // it is our head
		return
	}
	newHead := &ListItem{Value: i.Value}
	prevHead := l.head
	prevHead.Prev = newHead
	newHead.Next = prevHead
	newHead.Prev = nil
	l.head = newHead
	l.Remove(i)
}
