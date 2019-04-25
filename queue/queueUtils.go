package queue

import (
	"ds_go/linkedlist"
	"fmt"
)

//Queue ...
type Queue struct {
	Front, Rare *linkedlist.Node
}

//CreateQueue ...
func CreateQueue() *Queue {
	q := Queue{Front: nil, Rare: nil}
	return &q
}

//Enque ...
func (q *Queue) Enque(data interface{}) {
	newNode := linkedlist.CreateNode(data)
	if q.Front == nil {
		q.Front, q.Rare = newNode, newNode
	} else {
		q.Rare.Next = newNode
		q.Rare = newNode
	}
}

//Deque ...
func (q *Queue) Deque() *linkedlist.Node {
	if q.Rare == nil {
		return nil
	}
	tmp := q.Front
	for tmp.Next != q.Rare {
		tmp = tmp.Next
	}
	dequed := tmp.Next
	tmp.Next = nil
	q.Rare = tmp
	return dequed
}

//IsEmptyQ ...
func (q *Queue) IsEmptyQ() bool {
	if q.Front == nil {
		return true
	}
	return false
}

//GetFront ...
func (q *Queue) GetFront() *linkedlist.Node {
	if q.IsEmptyQ() {
		fmt.Println("Queue is Empty...")
		return nil
	}
	return q.Front
}
