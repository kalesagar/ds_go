package linkedlist

import "fmt"

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
func (n *Node) InsertAtLast(data interface{}) {
	newNode := CreateNode(data)
	tmp := n
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = newNode
}

//PrintList ...
func (n *Node) PrintList() {
	for n != nil {
		fmt.Printf("%v\n", n.Data)
		n = n.Next
	}
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
