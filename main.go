package main

import (
	"ds_go/linkedlist"
	"fmt"
)

func main() {
	fmt.Println("Started Main go routine....")
	llist := linkedlist.CreateNode(10)
	llist.InsertAtLast(15)
	llist = llist.InsertAtFirst(5)
	llist.InsertAtLast(20)
	llist = llist.InsertAtFirst(1)
	llist.InsertAtLast(30)
	llist.InsertAtLast(40)
	llist.PrintList()
	fmt.Println(llist.GetLength())
	fmt.Println(llist.GetMiddleOfList().Data)
	if kthlast := llist.FindKthLast(2); kthlast != nil {
		fmt.Println(kthlast.Data)
	}

}
