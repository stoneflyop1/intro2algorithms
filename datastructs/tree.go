package datastructs

import "fmt"

// Comparable compares two objects, used to implement collections sorting
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

// TreeSearchIterative searches value key in a binary search tree iteratively
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

// TreeSuccessor find the successor of a node in a binary search tree
func TreeSuccessor(node *BTreeNode) *BTreeNode {
	if node.right != nil {
		return TreeMinimum(node.right)
	}
	p := node.parent
	for p != nil && node == p.right {
		node = p
		p = p.parent
	}
	return p
}

// TreePredecessor find the predecessor of a node in a binary search tree
func TreePredecessor(node *BTreeNode) *BTreeNode {
	if node.left != nil {
		return TreeMaximum(node.left)
	}
	p := node.parent
	for p != nil && node == p.left {
		node = p
		p = p.parent
	}
	return p
}

// TreeInsert inserts element value v to a binary search tree,
// the inserted value must be Comparable interface
func TreeInsert(root *BTreeNode, v interface{}) *BTreeNode {
	v0, ok := v.(Comparable)
	if !ok {
		panic("must be comparable")
	}
	z := &BTreeNode{key: v}
	var p *BTreeNode
	r0 := root
	r := root
	for r != nil {
		p = r
		if v0.compare(r.key) < 0 {
			r = r.left
		} else {
			r = r.right
		}
	}
	z.parent = p
	if p == nil {
		r0 = z
	} else if v0.compare(p.key) < 0 {
		p.left = z
	} else {
		p.right = z
	}
	return r0
}

// TreeDelete deletes a node of binary search tree
func TreeDelete(root, z *BTreeNode) *BTreeNode {
	if z == nil {
		return root
	}
	if z.left == nil {
		root = transplant(root, z, z.right)
	} else if z.right == nil {
		root = transplant(root, z, z.left)
	} else {
		y := TreeMinimum(z.right)
		if y.parent != z {
			root = transplant(root, y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		root = transplant(root, z, y)
		y.left = z.left
		y.left.parent = y
	}
	return root
}

// transplant replaces one subtree as a child of its parent with another subtree.
// When transplant replaces the subtree rooted at node u with the subtree rooted at node v,
// node u's parent becomes node v's parent, and u's parent ends up having v as its appropriate child.
func transplant(root, u, v *BTreeNode) *BTreeNode {
	if u.parent == nil {
		root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
	return root
}

// TreeNode general treenode with each node can have one left child and
// many rightSiblings which is started with rightSibling
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
