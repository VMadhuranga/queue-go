package main

type node struct {
	data any
	next *node
}

type queue struct {
	head, tail *node
	size       int
}

func (q *queue) push(data any) {
	n := &node{
		data: data,
		next: nil,
	}

	if q.head == nil {
		q.head = n
	} else {
		q.tail.next = n
	}

	q.tail = n
	q.size++
}

func (q *queue) pop() any {
	n := &node{}

	if q.head != nil {
		n = q.head
		q.head = q.head.next
		q.size--
	}

	return n.data
}

func (q *queue) peek() any {
	if q.head == nil {
		return nil
	}

	return q.head.data
}

func (q *queue) getSize() int {
	return q.size
}
