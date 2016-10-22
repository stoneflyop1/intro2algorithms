package datastructs

import "fmt"

//todo

// BTreeNode Binary TreeNode
type BTreeNode struct {
	key    interface{}
	parent *BTreeNode
	left   *BTreeNode
	right  *BTreeNode
}

func PrintBinaryTree(root *BTreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.key)
	if root.left != nil {
		PrintBinaryTree(root.left)
	}
	if root.right != nil {
		PrintBinaryTree(root.right)
	}
}

func PrintBinaryTree2(root *BTreeNode) {
	if root == nil {
		return
	}
	r := root.right
	for root != nil {
		fmt.Println(root.key)
		root = root.left
	}
	if r != nil {
		fmt.Println(r.key)
		r = r.right
	}
}

// TreeNode
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
