package list

import "fmt"

// Element s
type Element struct {
	data int
	next *Element
}

// List s
type List struct {
	head *Element
	size int
}

// New s
func New(items []int) *List {
	l := &List{}
	for _, e := range items {
		l.Push(e)
	}
	return l
}

// Size s
func (l *List) Size() int {
	return l.size
}

// Push s
func (l *List) Push(e int) {
	var p, last *Element

	elem := &Element{
		data: e,
		next: nil,
	}

	if l.size == 0 {
		l.head = elem
		l.size++
		return
	}

	for p = l.head; p != nil; p = p.next {
		last = p
	}

	last.next = elem
	l.size++
}

// Pop s
func (l *List) Pop() (int, error) {

	if l.size == 0 {
		return 0, fmt.Errorf("empty list")
	}

	if l.size == 1 {
		v := l.head.data
		l.head = nil
		l.size--
		return v, nil
	}

	var p, last *Element
	for p = l.head; p.next != nil; p = p.next {
		last = p
	}

	value := last.next.data
	last.next = nil
	l.size--

	return value, nil
}

// Array s
func (l *List) Array() []int {
	items := make([]int, l.size)

	for p, i := l.head, 0; p != nil; p, i = p.next, i+1 {
		items[i] = p.data
	}

	return items
}

// Reverse s
func (l *List) Reverse() *List {
	arr := l.Array()
	l = nil

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return New(arr)
}
