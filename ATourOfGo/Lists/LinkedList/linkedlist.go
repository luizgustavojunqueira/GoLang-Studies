package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	next *List[T]
	val  T
}

// NewList initialize a new list
func NewList[T comparable]() *List[T] {
	return &List[T]{}
}

// Append adds a new value to the end of the list.
func (l *List[T]) Append(val T) {
	p := l
	var zero T
	if p.val == zero {
		p.val = val
		return
	}
	for p.next != nil {
		p = p.next
	}
	p.next = &List[T]{val: val}
}

// Print prints the list to the console.
func (l *List[T]) Print() {
	p := l
	for p != nil {
		fmt.Printf("%v -> ", p.val)
		p = p.next
	}
	fmt.Println()
}

// Delete remove the first ocurrence of given value
func (l *List[T]) Delete(val T) (index, deleted int) {
	index = 0
	p := l
	for p.next != nil {
		if p.next.val == val {
			p.next = p.next.next
			return index + 1, 1
		}
		p = p.next
		index += 1
	}
	return index + 1, 0
}

func (l *List[T]) DeleteAll(val T) (count int) {
	p := l
	for p.next != nil {
		if p.next.val == val {
			p.next = p.next.next
			count += 1
		}
		if p.next.val != val {
			p = p.next
		}
	}
	return
}

func (l *List[T]) Search(val T) (index int) {
	p := l
	for p != nil {
		if p.val == val {
			return index
		}
		p = p.next
		index += 1
	}

	return
}

func main() {
	l := NewList[int]()

	l.Append(1)
	l.Append(2)
	l.Append(2)
	l.Append(1)
	l.Append(1)
	l.Append(3)
	l.Append(1)
	l.Print()
	index, deleted := l.Delete(5)
	fmt.Printf("Index: %d, Deleted: %d\n", index, deleted)
	l.Print()

	index = l.Search(2)
	fmt.Printf("Value: %v is at index %v\n", 2, index)

	count := l.DeleteAll(2)
	fmt.Printf("Delete %v values\n", count)
	l.Print()

	index = l.Search(2)
	fmt.Printf("Value: %v is at index %v\n", 2, index)
}
