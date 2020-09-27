package main

import "fmt"

// Node is a struct
type Node struct {
	left  *Node
	data  interface{}
	right *Node
}

// NewNode is a function which is used to create a new node
func NewNode(le, rt *Node) *Node {
	return &Node{le, nil, rt}
}

// SetData is a method of struct Node, which is used to set data to Node
func (node *Node) SetData(data interface{}) {
	node.data = data
}

func main() {

	root := NewNode(nil, nil)
	root.SetData("RootNode")

	le := NewNode(nil, nil)
	le.SetData("LeftNode")
	rt := NewNode(nil, nil)
	rt.SetData("RightNode")

	root.left = le
	root.right = rt

	fmt.Printf("%v\n", root)
}
