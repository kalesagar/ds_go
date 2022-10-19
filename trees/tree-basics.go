package trees

import (
	"errors"
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type Queue struct {
	data []*Node
}

func (Q *Queue) Push(d *Node) {
	Q.data = append(Q.data, d)
	return
}

func (Q *Queue) Pop() (*Node, error) {
	n := len(Q.data)
	if n == 0 {
		return nil, errors.New("Can't pop elements.. Queue is empty..")
	}
	value := Q.data[0]
	if n == 1 {
		Q.data = []*Node{}
	} else {
		Q.data = Q.data[1:n]
	}
	return value, nil
}

func (Q *Queue) Front() *Node {
	n := len(Q.data)
	if n == 0 {
		return nil
	}
	return Q.data[0]
}

func (Q *Queue) IsEmpty() bool {
	n := len(Q.data)
	if n == 0 {
		return true
	}
	return false
}

func InsertNode(root *Node, d int) *Node {
	if root == nil {
		root = CreateNode(d)
		return root
	}
	q := Queue{}
	q.Push(root)
	for !q.IsEmpty() {
		tmp, _ := q.Pop()
		if tmp.left == nil {
			tmp.left = CreateNode(d)
			return root
		} else {
			q.Push(tmp.left)
		}
		if tmp.right == nil {
			tmp.right = CreateNode(d)
			return root
		} else {
			q.Push(tmp.right)
		}
	}
	return root
}
func CreateNode(data int) *Node {
	return &Node{data: data, left: nil, right: nil}
}

func CreateBinaryTree() *Node {
	root := CreateNode(1)
	root.left = CreateNode(2)
	root.right = CreateNode(3)
	root.left.left = CreateNode(4)
	root.left.right = CreateNode(5)
	//root.right.left = CreateNode(6)
	root.right.right = CreateNode(7)
	return root
}

func InorderTraversal(root *Node) {
	if root == nil {
		return
	}

	InorderTraversal(root.left)
	fmt.Printf("%d ", root.data)
	InorderTraversal(root.right)
}

func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.data)
	PreOrderTraversal(root.left)
	PreOrderTraversal(root.right)
}

func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.left)
	PostOrderTraversal(root.right)
	fmt.Printf("%d ", root.data)
}

func DeleteNodeFromTree(root *Node, d int) *Node {
	var q Queue
	if root.data == d {
		rNode := root.right
		rNode.left = root.left
		return rNode
	}
	q.Push(root)
	for !q.IsEmpty() {

	}
	return root
}
