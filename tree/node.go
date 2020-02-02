package node

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
	//in c++,it will make a crash,but in golang,it's ok
}

func (node TreeNode) PrintNode() { //receiver
	fmt.Print(node.Value)
}

func (node *TreeNode) SetValue(value int) {
	node.Value = value
}

func (node *TreeNode) Travese() {
	if nil == node {
		return
	}
	node.Left.Travese()
	node.PrintNode()
	node.Right.Travese()
}
