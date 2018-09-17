package linkedlist

import "testing"

var linkedList1 LinkedList
var linkedList2 LinkedList
var linkedList3 LinkedList

func TestMergeTwoSortedLinkedList(t *testing.T) {
	linkedList1.Append(1)
	linkedList1.Append(5)
	linkedList1.Append(10)
	linkedList2.Append(2)
	linkedList2.Append(7)
	linkedList2.Append(21)
	linkedList2.Append(23)
	linkedList2.Append(23)
	linkedList3 = MergeTwoSortedLinkedList(linkedList1, linkedList2)
	if linkedList3.IsEmpty() {
		t.Errorf("list should not be empty")
	}
	if size := linkedList3.Size(); size != 8 {
		t.Errorf("wrong count, expected 8 and got %d", size)
	}
}
