package main

import (
	"fmt"

	node "github.com/sniancmk/learnGo/tree"
)

func main() {
	var root node.TreeNode
	root = node.TreeNode{Value: 3}
	root.Left = &node.TreeNode{}
	root.Right = &node.TreeNode{
		Value: 5,
		Left:  nil,
		Right: nil,
	}
	root.Right.Left = new(node.TreeNode)
	root.Left.Right = node.CreateNode(2)

	nodes := []node.TreeNode{
		{Value: 3},
		{},
		{6, nil, nil},
	}
	fmt.Println(nodes)
	root.Right.Left.SetValue(4)
	root.Right.Left.PrintNode()
	root.PrintNode()
	root.Travese()
}
