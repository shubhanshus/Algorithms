package linkedlist

// RemoveDuplicates removes duplicates from LinkedList
func (linkedList *LinkedList) RemoveDuplicates() {
	linkedList.lock.Lock()
	defer linkedList.lock.Unlock()

	node := linkedList.head

	for node != nil && node.next != nil {
		nextNode := node
		for nextNode.next != nil {
			if node.content == nextNode.next.content {
				nextNode.next = nextNode.next.next
				linkedList.size--
			} else {
				nextNode = nextNode.next
			}
		}
		node = node.next
	}
}

//DeleteNode deletes the node "node" from the LinkedList
// This does not update the size of the linkedlist which it should do
func (node *Node) DeleteNode() {
	for node != nil && node.next != nil {
		node.content = node.next.content
		node.next = node.next.next
	}
}

// Partition function for CCCI Q 2.4
func (linkedList *LinkedList) Partition(node int) {
	var currentNode *Node
	currentNode = linkedList.head
	for currentNode != nil && currentNode.next != nil {

	}
}
