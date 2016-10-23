package datastructs

import "fmt"

type Comparable interface {
	compare(v interface{}) int
}

// BTreeNode Binary TreeNode
type BTreeNode struct {
	key    interface{}
	parent *BTreeNode
	left   *BTreeNode
	right  *BTreeNode
}

// PrintBinaryTree preorder recursion
func PrintBinaryTree(root *BTreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.key, "\t")
	if root.left != nil {
		PrintBinaryTree(root.left)
	}
	if root.right != nil {
		PrintBinaryTree(root.right)
	}
}

// PrintBinaryTree2 inorder tranverse
// see http://www.geeksforgeeks.org/inorder-tree-traversal-without-recursion/
// 1. Create an empty stack S.
// 2. Initialize current node as root
// 3. Push current node to S and set current = current.left until current = nil
// 4. If current is nil and stack is not empty then
//      1) Pop the top item from stack.
//      2) Print the popped item, set current = poppedItem.right
//      3) Go to step 3.
// 5. If current is nil and stack is empty then we are done
func PrintBinaryTree2(root *BTreeNode) {
	if root == nil {
		return
	}
	stack := NewListStack()
	for root != nil {
		stack.Push(root)
		root = root.left
	}
	for !stack.IsEmpty() {
		p := stack.Pop().(*BTreeNode) //type assertion, see https://tour.golang.org/methods/15
		fmt.Print(p.key, "\t")
		root = p.right
		for root != nil {
			stack.Push(root)
			root = root.left
		}
	}
}

// TreeSearch search in a binary search tree recursively
func TreeSearch(tree *BTreeNode, key interface{}) *BTreeNode {

	c1, ok1 := key.(Comparable)
	if !ok1 {
		panic("key must implement Comparable interafce")
	}
	c := c1.compare(tree.key)
	if tree == nil || c == 0 {
		return tree
	}
	if c < 0 {
		return TreeSearch(tree.left, key)
	} else {
		return TreeSearch(tree.right, key)
	}
}

// TreeSearchIterative search in a binary search tree iteratively
func TreeSearchIterative(tree *BTreeNode, key interface{}) *BTreeNode {
	c1, ok1 := key.(Comparable)
	if !ok1 {
		panic("key must implement Comparable interafce")
	}
	c := c1.compare(tree.key)
	for tree != nil && c != 0 {
		if c < 0 {
			tree = tree.left
		} else {
			tree = tree.right
		}
		c = c1.compare(tree.key)
	}
	return tree
}

// TreeMinimum find the minimum element in a binary search tree
func TreeMinimum(tree *BTreeNode) *BTreeNode {
	if tree == nil {
		return nil
	}
	for tree.left != nil {
		tree = tree.left
	}
	return tree
}

// TreeMaximum find the maximum element in a binary search tree
func TreeMaximum(tree *BTreeNode) *BTreeNode {
	if tree == nil {
		return nil
	}
	for tree.right != nil {
		tree = tree.right
	}
	return tree
}

// TreeNode
type TreeNode struct {
	key          interface{}
	parent       *TreeNode
	leftChild    *TreeNode
	rightSibling *TreeNode
}

//PrintTree  preorder recursion
func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.key, "\t")
	if root.leftChild != nil {
		PrintTree(root.leftChild)
	}
	if root.rightSibling != nil {
		PrintTree(root.rightSibling)
	}
}
