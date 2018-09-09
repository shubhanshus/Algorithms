package linkedlist

import (
	"fmt"
	"testing"
)

var linkedList LinkedList

func TestAppend(t *testing.T) {
	if !linkedList.IsEmpty() {
		t.Errorf("list should be empty")
	}

	linkedList.Append("first")
	if linkedList.IsEmpty() {
		t.Errorf("list should not be empty")
	}

	if size := linkedList.Size(); size != 1 {
		t.Errorf("wrong count, expected 1 and got %d", size)
	}

	linkedList.Append("second")
	linkedList.Append("third")

	if size := linkedList.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestRemoveAt(t *testing.T) {
	_, err := linkedList.RemoveAt(1)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if size := linkedList.Size(); size != 2 {
		t.Errorf("wrong count, expected 2 and got %d", size)
	}
}

func TestInsert(t *testing.T) {
	//test inserting in the middle
	err := linkedList.Insert(2, "second2")
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if size := linkedList.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}

	//test inserting at head position
	err = linkedList.Insert(0, "zero")
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
}

func TestIndexOf(t *testing.T) {
	if i := linkedList.IndexOf("zero"); i != 0 {
		t.Errorf("expected position 0 but got %d", i)
	}
	if i := linkedList.IndexOf("first"); i != 1 {
		t.Errorf("expected position 1 but got %d", i)
	}
	if i := linkedList.IndexOf("second2"); i != 2 {
		t.Errorf("expected position 2 but got %d", i)
	}
	if i := linkedList.IndexOf("third"); i != 3 {
		t.Errorf("expected position 3 but got %d", i)
	}
}

func TestHead(t *testing.T) {
	h := linkedList.Head()
	if "zero" != fmt.Sprint(h.content) {
		t.Errorf("Expected `zero` but got %s", fmt.Sprint(h.content))
	}
}
