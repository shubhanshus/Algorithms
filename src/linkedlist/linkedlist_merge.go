package linkedlist

//MergeTwoSortedLinkedList function is clasic problem of merging two sorted LinkedList
//This function returns the sorted LinkedList back which contains data from both the linkedlists
func MergeTwoSortedLinkedList(linkedlist1, linkedlist2 LinkedList) LinkedList {
	var finalLinkedList LinkedList
	nodes := make([]*Node, 2)
	nodes[0] = linkedlist1.head
	nodes[1] = linkedlist2.head

	for nodes[0] != nil || nodes[1] != nil {
		if nodes[0] != nil && nodes[1] != nil {
			if nodes[0].content.(int) < nodes[1].content.(int) {
				finalLinkedList.Append(nodes[0].content)
				nodes[0] = nodes[0].next
			} else {
				finalLinkedList.Append(nodes[1].content)
				nodes[1] = nodes[1].next
			}
		} else if nodes[0] != nil {
			finalLinkedList.Append(nodes[0].content)
			nodes[0] = nodes[0].next
		} else {
			finalLinkedList.Append(nodes[1].content)
			nodes[1] = nodes[1].next
		}
	}
	return finalLinkedList
}
