package linkedlist

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item to be stored in LinkedList
type Item generic.Type

// Node of a LinkedList
type Node struct {
	content Item
	next    *Node
}

// LinkedList is the linkedlist of Nodes
type LinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

//Append function is for appending an item into LinkedList
func (linkedList *LinkedList) Append(item Item) {
	linkedList.lock.Lock()
	defer linkedList.lock.Unlock()
	node := Node{item, nil}
	if linkedList.head == nil {
		linkedList.head = &node
	} else {
		last := linkedList.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	linkedList.size++

}

// IsEmpty returns true if LinkedList is empty
func (linkedList *LinkedList) IsEmpty() bool {
	linkedList.lock.RLock()
	defer linkedList.lock.RUnlock()
	if linkedList.head == nil {

		return true
	}
	return false
}

// Size returns the size of the LinkedList
func (linkedList *LinkedList) Size() int {
	linkedList.lock.RLock()
	defer linkedList.lock.RUnlock()
	if linkedList.head == nil {

		return 0
	}
	return linkedList.size
}

//IndexOf returns the index of item in the LinkedList
func (linkedList *LinkedList) IndexOf(item Item) int {
	linkedList.lock.Lock()
	defer linkedList.lock.Unlock()
	if linkedList.head == nil {

		return 0
	}
	pos := 0
	nextNode := linkedList.head
	for {
		if nextNode.content == item {

			return pos
		}
		if nextNode.next == nil {

			return 0
		}
		nextNode = nextNode.next
		pos++
	}
}

//RemoveAt removes the item from the LinkedList position i
func (linkedList *LinkedList) RemoveAt(i int) (Item, error) {
	linkedList.lock.Lock()
	defer linkedList.lock.Unlock()
	if i < 0 || i > linkedList.size {

		return nil, fmt.Errorf("Index out of bounds")
	}
	node := linkedList.head

	j := 0
	for j < i-1 {
		j++
		node = node.next
	}
	remove := node.next
	node.next = remove.next
	linkedList.size = linkedList.size - 1
	return &remove.content, nil

}

// Insert adds a new node in the LinkedList at the position specified
func (linkedList *LinkedList) Insert(i int, item Item) error {
	linkedList.lock.Lock()
	defer linkedList.lock.Unlock()
	if i < 0 || i > linkedList.size {

		return fmt.Errorf("Index out of bounds")
	}
	newNode := Node{item, nil}
	if i == 0 {
		newNode.next = linkedList.head
		linkedList.head = &newNode
		return nil
	}
	node := linkedList.head
	j := 0
	for j < i-2 {
		j++
		node = node.next
	}
	newNode.next = node.next
	node.next = &newNode
	linkedList.size++
	return nil
}

// Head returns the head of the LinkedList
func (linkedList *LinkedList) Head() *Node {
	linkedList.lock.RLock()
	defer linkedList.lock.RUnlock()
	return linkedList.head
}

func (linkedList *LinkedList) String() {
	linkedList.lock.RLock()
	defer linkedList.lock.RUnlock()
	node := linkedList.head
	j := 0
	for {
		if node == nil {

			break
		}
		j++
		fmt.Print(node.content)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()

}
