package queue

type MyCircularQueue struct {
	size int
	front int
	rear int
	elements []int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	myQueue:= MyCircularQueue{
		size:k+1,
		front:0,
		rear:0,
		elements:make([]int,k+1),
	}
	return myQueue
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull(){
		return false
	}
	this.elements[this.rear]=value
	this.rear=(this.rear+1)%this.size
	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty(){
		return false
	}
	//this.elements[this.front]=nil
	this.front=(this.front+1)%this.size
	return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty(){
		return -1
	}
	pos:= (this.front+this.size)%this.size
	return this.elements[pos]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {

	if this.IsEmpty(){
		return -1
	}
	pos:= (this.rear+this.size-1)%this.size
	return this.elements[pos]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if (this.rear)%this.size == this.front%this.size {
		return true
	}
	return false
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if (this.rear+1)%this.size == this.front%this.size {
		return true
	}
	return false
}
