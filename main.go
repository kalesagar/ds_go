package main

import (
	//"ds_go/linkedlist"
	"ds_go/programs"
	//"ds_go/queue"
	//"ds_go/programs"
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
	// llist = llist.ReverseLinkedListNonRecursive()
	// llist.PrintList()
	// llist = llist.DeleteNodeFromLinkedList(15)
	// llist.PrintList()
	// llist.PrintLinnkedListInReverseOrder()

	// llist1 := linkedlist.CreateNode(11)
	// llist1.InsertAtLast(17)
	// llist1 = llist1.InsertAtFirst(5)
	// llist1.InsertAtLast(22)
	// llist1 = llist1.InsertAtFirst(1)
	// llist1 = llist1.InsertAtFirst(1)
	// llist1.InsertAtLast(33)
	// llist1.InsertAtLast(5)
	// llist1.InsertAtLast(45)
	// llist1.InsertAtLast(55)
	// llist1.InsertAtLast(55)
	// llist1.PrintList()
	//llist1 = llist1.DeleteDuplicateNodesFromUnSortedLinkedList()
	// llist1 = llist1.DeleteDuplicateNodesFromUnSortedLinkedList()
	// llist1.PrintList()
	// fmt.Println(linkedlist.CompareTwoLinkedList(llist, llist1))
	// fmt.Println()
	// newLL := linkedlist.MergeSortedLinkedLists(llist, llist1)
	// newLL.PrintList()
	// newLL.InsertAtLast(55)
	// newLL.PrintList()
	// newLL.DeleteDuplicateNodesFromSortedLinkedList()
	// newLL.PrintList()
	// *** create cycle in linked list ***
	// tmp := newLL.InsertAtLast(56)
	// tmp.Next = newLL.Next
	// fmt.Println(newLL.CycleDetectionInLinkedList())
	// *** removing cycle from linked list ***
	// tmp.Next = nil
	// fmt.Println(newLL.CycleDetectionInLinkedList())
	// ***
	// fmt.Println(llist.GetLength())
	// fmt.Println(llist.GetMiddleOfList().Data)
	// if kthlast := llist.FindKthLast(2); kthlast != nil {
	// 	fmt.Println(kthlast.Data)
	// }
	//-------------Queue---------------
	// q := queue.CreateQueue()
	// fmt.Println(q.IsEmptyQ())
	// q.Enque(15)
	// q.Enque(20)
	// q.Enque(30)
	// q.Front.PrintList()
	// q.Enque(40)
	// q.Enque(50)
	// if dequed := q.Deque(); dequed != nil {
	// 	fmt.Println(dequed.Data)
	// }
	// if dequed := q.Deque(); dequed != nil {
	// 	fmt.Println(dequed.Data)
	// }
	// front := q.GetFront()
	// if front != nil {
	// 	fmt.Println(front.Data)
	// }
	//---------program FindLargestNumber-----------
	// largeNo := programs.LargeNumber([]int{78,66,57})
	// fmt.Println(largeNo)

	// programs.PairsOfSum([]int{10, 20, 5, 15, 7, 8}, 25)

	//------ print even and odd numbers using two go routines

	//programs.PrintOddNumbersUsingGoRoutines()

	//----- Check string is pangram ----
	fmt.Println(programs.IsPangram("The quick brown fox jumps over the lazy dog")) //true
	fmt.Println(programs.IsPangram("The quick brown fox jumps over the dog")) //false
	//--- MissingCharacterToPangramString ---
	programs.MissingCharacterToPangramString("The quick brown fox jumps over the lazy dog")
	programs.MissingCharacterToPangramString("The quick brown fox jumps over the dog")
}
