package datastructs

import "fmt"

//todo

type TreeNode struct {
	key          interface{}
	parent       *TreeNode
	leftChild    *TreeNode
	rightSibling *TreeNode
}

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.key)
	if root.leftChild != nil {
		PrintTree(root.leftChild)
	}
	if root.rightSibling != nil {
		PrintTree(root.rightSibling)
	}
}
