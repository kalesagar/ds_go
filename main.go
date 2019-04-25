package main

import (
	//"ds_go/linkedlist"
	"ds_go/queue"
	"fmt"
)

func main() {
	fmt.Println("Started Main go routine....")
	//-------------------Linked List------------------------
	// llist := linkedlist.CreateNode(10)
	// llist.InsertAtLast(15)
	// llist = llist.InsertAtFirst(5)
	// llist.InsertAtLast(20)
	// llist = llist.InsertAtFirst(1)
	// llist.InsertAtLast(30)
	// llist.InsertAtLast(40)
	// llist.PrintList()
	// fmt.Println(llist.GetLength())
	// fmt.Println(llist.GetMiddleOfList().Data)
	// if kthlast := llist.FindKthLast(2); kthlast != nil {
	// 	fmt.Println(kthlast.Data)
	// }
	//-------------Queue---------------
	q := queue.CreateQueue()
	fmt.Println(q.IsEmptyQ())
	q.Enque(15)
	q.Enque(20)
	q.Enque(30)
	q.Front.PrintList()
	q.Enque(40)
	q.Enque(50)
	if dequed := q.Deque(); dequed != nil {
		fmt.Println(dequed.Data)
	}
	if dequed := q.Deque(); dequed != nil {
		fmt.Println(dequed.Data)
	}
	front := q.GetFront()
	if front != nil {
		fmt.Println(front.Data)
	}
}
