package linkedlist

import (
	"fmt"
)

//Node ...
type Node struct {
	Data interface{}
	Next *Node
}

//CreateNode ...
func CreateNode(data interface{}) *Node {
	newNode := Node{Data: data, Next: nil}
	return &newNode
}

//InsertAtLast ...
func (n *Node) InsertAtLast(data interface{}) *Node{
	newNode := CreateNode(data)
	tmp := n
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = newNode
	return newNode
}

//PrintList ...
func (n *Node) PrintList() {
	for n != nil {
		fmt.Printf("%v --> ", n.Data)
		n = n.Next
	}
	fmt.Print("nil \n")
}

//InsertAtFirst ...
func (n *Node) InsertAtFirst(data interface{}) *Node {
	newNode := CreateNode(data)
	newNode.Next = n
	return newNode
}

//GetLength ...
func (n *Node) GetLength() int {
	tmp := n
	length := 0
	for tmp != nil {
		length++
		tmp = tmp.Next
	}
	return length
}

//GetMiddleOfList ...
func (n *Node) GetMiddleOfList() *Node {
	first, second := n, n
	for second != nil {
		if second.Next != nil {
			second = second.Next.Next
			first = first.Next
		} else {
			break
		}
	}
	return first
}

//FindKthLast ...
func (n *Node) FindKthLast(k int) *Node {
	first, second := n, n
	i := 0
	for i < k {
		if second != nil {
			second = second.Next
			i++
		} else {
			fmt.Println("ListsLength Is less than K...please select K less than List length...")
			return nil
		}
	}
	for second != nil {
		second = second.Next
		first = first.Next
	}
	return first
}

//ReverseLinkedListNonRecursive ...
func (head *Node) ReverseLinkedListNonRecursive() *Node {
	var prev *Node
	current := head
	for current != nil {
		nextNode := current.Next
		current.Next = prev
		prev = current
		current = nextNode
	}
	return prev
}

//DeleteNodeFromLinkedList ...
func (head *Node) DeleteNodeFromLinkedList(data int) *Node {
	current := head
	if current.Data == data {
		return current.Next
	}
	prev := current
	current = current.Next
	for current != nil {
		nextNode := current.Next
		if current.Data == data {
			prev.Next = nextNode
			return head
		}
		prev = current
		current = current.Next
	}
	return head
}

// PrintLinnkedListInReverseOrder ...
func (head *Node) PrintLinnkedListInReverseOrder() {
	if head == nil {
		return
	}
	if head.Next == nil {
		fmt.Printf("nil --> %v -->", head.Data)
		return
	}
	head.Next.PrintLinnkedListInReverseOrder()
	fmt.Printf("%v -->", head.Data)
}

//CompareTwoLinkedList ...
func CompareTwoLinkedList(head1 *Node, head2 *Node) bool {
	if head1 == nil && head2 == nil {
		return true
	}
	for head1 != nil && head2 != nil {
		if head1.Data != head2.Data {
			return false
		}
		head1 = head1.Next
		head2 = head2.Next
	}
	if head1 == nil && head2 == nil {
		return true
	}
	return false
}

//MergeSortedLinkedLists ...
func MergeSortedLinkedLists(head1 *Node, head2 *Node) *Node {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	var head *Node
	if head1.Data.(int) < head2.Data.(int) {
		head = head1
		head1 = head1.Next
	} else {
		head = head2
		head2 = head2.Next
	}
	realHead := head
	for head1 != nil && head2 != nil {
		if head1.Data.(int) < head2.Data.(int) {
			head.Next = head1
			head = head1
			head1 = head1.Next
		} else {
			head.Next = head2
			head = head2
			head2 = head2.Next
		}
	}
	if head1 == nil {
		head.Next = head2
	}
	if head2 == nil {
		head.Next = head1
	}
	return realHead
}

//DeleteDuplicateNodesFromSortedLinkedList ...
func (head *Node) DeleteDuplicateNodesFromSortedLinkedList() *Node {
	current := head
	for current != nil {
		if current.Data == current.Next.Data {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

//CycleDetectionInLinkedList ...
func (head *Node) CycleDetectionInLinkedList() bool{
	if head ==nil || head.Next == nil{
		return false
	}
	first := head
	second := head.Next.Next
	for first != nil && second != nil{
		if first == second{
			return true
		}
		first = first.Next
		if second.Next != nil{
			second = second.Next.Next
		}
	}
	return false
}


