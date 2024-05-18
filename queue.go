package main

import "fmt"

// Singly linked list implementation

type node struct {
	data int
	next *node
}

type singlyLinkedList struct {
	head *node
	tail *node
	size int
}

type queue struct {
	singlyLinkedList
}

func (q *queue) enqueue(data int) {
	node := &node{
		data: data,
		next: nil,
	}

	if q.head == nil {
		q.head = node
	} else {
		q.tail.next = node
	}

	q.tail = node
	q.size++
}

func (q *queue) dequeue() {
	if q.head == nil {
		fmt.Println("Queue is empty")
	} else {
		q.head = q.head.next
		q.size--
	}
}

func (q queue) isEmpty() {
	if q.head == nil {
		fmt.Println("Queue is empty")
	} else {
		fmt.Println("Queue is not empty")
	}
}

func (q queue) printFirst() {
	if q.head == nil {
		fmt.Println("Queue is empty")
	} else {
		fmt.Println(q.head.data)
	}
}

func (q queue) printSize() {
	fmt.Println(q.size)
}

func main() {
	q := queue{}
	q.isEmpty()         // "Queue is empty"
	q.printFirst()      // "Queue is empty"
	q.dequeue()         // "Queue is empty"
	q.printSize()       // 0
	fmt.Println(q.head) // nil
	fmt.Println(q.tail) // nil

	q.enqueue(1)        // 1 -> nil
	q.enqueue(3)        // 1 -> 3 -> nil
	q.enqueue(5)        // 1 -> 3 -> 5 -> nil
	q.isEmpty()         // "Queue is not empty"
	q.printFirst()      // 1
	q.printSize()       // 3
	fmt.Println(q.head) // 1
	fmt.Println(q.tail) // 5

	q.dequeue()         // 3 -> 5 -> nil
	q.isEmpty()         // "Queue is not empty"
	q.printFirst()      // 3
	q.printSize()       // 2
	fmt.Println(q.head) // 3
	fmt.Println(q.tail) // 5
}

// Slices implementation

// type queue struct {
// 	elements []int
// }

// func (q *queue) enqueue(data int) {
// 	q.elements = append(q.elements, data)
// }

// func (q *queue) dequeue() {
// 	if len(q.elements) <= 0 {
// 		fmt.Println("Queue is empty")
// 	} else {
// 		q.elements = q.elements[1:]
// 	}
// }

// func (q queue) isEmpty() {
// 	if len(q.elements) <= 0 {
// 		fmt.Println("Queue is empty")
// 	} else {
// 		fmt.Println("Queue is not empty")
// 	}
// }

// func (q queue) printFirst() {
// 	if len(q.elements) <= 0 {
// 		fmt.Println("Queue is empty")
// 	} else {
// 		fmt.Println(q.elements[0])
// 	}
// }

// func (q queue) printSize() {
// 	fmt.Println(len(q.elements))
// }

// func main() {
// 	q := queue{}
// 	q.isEmpty()    // "Queue is empty"
// 	q.printFirst() // "Queue is empty"
// 	q.dequeue()    // "Queue is empty"
// 	q.printSize()  // 0

// 	q.enqueue(1)   // [1]
// 	q.enqueue(3)   // [1, 3]
// 	q.enqueue(5)   // [1, 3, 5]
// 	q.isEmpty()    // "Queue is not empty"
// 	q.printFirst() // 1
// 	q.printSize()  // 3

// 	q.dequeue()    // [3, 5]
// 	q.isEmpty()    // "Queue is not empty"
// 	q.printFirst() // 3
// 	q.printSize()  // 2
// }
