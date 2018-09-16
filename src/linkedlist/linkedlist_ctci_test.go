package linkedlist

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	linkedList.Append(4)
	linkedList.Append(3)
	linkedList.RemoveDuplicates()
	if size := linkedList.Size(); size != 5 {
		t.Errorf("wrong count, expected 5 and got %d", size)
	}
}
