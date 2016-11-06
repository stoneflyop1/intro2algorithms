package datastructs

import "fmt"

const (
	RED, BLACK = 0, 1
)

type RbTreeNode struct {
	parent, left, right *RbTreeNode
	color               int
	key                 interface{}
}

type RbTree struct {
	root *RbTreeNode
}

func RbTreeHeight(root *RbTreeNode) int {
	if root == nil {
		return -1
	}
	left := RbTreeHeight(root.left)
	right := RbTreeHeight(root.right)
	if left < right {
		return right + 1
	} else {
		return left + 1
	}
}

// PrintBinaryTree inorder recursion
func PrintRbBinaryTree(root *RbTreeNode) {
	if root == nil {
		return
	}
	if root.left != nil {
		PrintRbBinaryTree(root.left)
	}
	fmt.Print(root.key, "\t")
	if root.right != nil {
		PrintRbBinaryTree(root.right)
	}
}

func RbTreeSearch(x *RbTreeNode, k interface{}) *RbTreeNode {
	k1, ok := k.(Comparable)
	if !ok {
		panic("key must be omparable")
	}
	cv := k1.compare(x.key)
	for x != nil && cv != 0 {
		cv = k1.compare(x.key)
		if cv < 0 {
			x = x.left
		} else if cv > 0 {
			x = x.right
		} else {
			return x
		}
	}
	return x
}

func RbTreeMinimum(x *RbTreeNode) *RbTreeNode {
	if x == nil {
		return nil
	}
	if x.left != nil {
		return RbTreeMinimum(x.left)
	}
	return x
}

func RbTreeMaximum(x *RbTreeNode) *RbTreeNode {
	if x == nil {
		return nil
	}
	if x.right != nil {
		return RbTreeMaximum(x.right)
	}
	return x
}

func RbTreeSuccessor(x *RbTreeNode) *RbTreeNode {
	if x == nil {
		return nil
	}
	if x.right != nil {
		return RbTreeMinimum(x.right)
	}
	p := x.parent
	for p != nil && p.right == x {
		x = p
		p = p.parent
	}
	return p
}

func RbTreePredecessor(x *RbTreeNode) *RbTreeNode {
	if x == nil {
		return nil
	}
	if x.left != nil {
		return RbTreeMaximum(x.left)
	}
	p := x.parent
	for p != nil && p.left == x {
		x = p
		p = p.parent
	}
	return p
}

func (T *RbTree) leftRotate(x *RbTreeNode) {
	y := x.right
	//x的右子树转换为y的左子树
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	//转换x的parent给y做parent
	y.parent = x.parent
	if x.parent == nil {
		T.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	// y作为x的父节点
	y.left = x
	x.parent = y
}

func (T *RbTree) rightRotate(y *RbTreeNode) {
	x := y.left
	//y的左子树转换为x的右子树
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	//转换y的parent给x做parent
	x.parent = y.parent
	if y.parent == nil {
		T.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}
	//x作为y的父节点
	x.right = y
	y.parent = x
}

func (T *RbTree) RbTreeInsert(v interface{}) {
	v0, ok := v.(Comparable)
	if !ok {
		panic("must be comparable")
	}
	z := &RbTreeNode{key: v, color: RED} // set red to satisfy black height property
	x := T.root
	y := T.root // 记录不为nil的x
	for x != nil {
		y = x
		c := v0.compare(x.key)
		if c < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {
		T.root = z
	} else if v0.compare(y.key) < 0 {
		y.left = z
	} else {
		y.right = z
	}
	//z.left = nil
	//z.right = nil
	//z.color = RED
	fmt.Println("insert value: ", v, z.parent)
	T.rbTreeInsertFixup(z)
}

func (T *RbTree) rbTreeInsertFixup(z *RbTreeNode) {
	if z.parent == nil {
		return
	}
	for z.parent != nil && z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y != nil && y.color == RED { // case 1
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right { // case 2
					z = z.parent
					T.leftRotate(z)
				}
				z.parent.color = BLACK         // case 3
				z.parent.parent.color = RED    // case 3
				T.rightRotate(z.parent.parent) // case 3
			}
		} else { // z.parent == z.parent.parent.right
			y := z.parent.parent.left
			if y != nil && y.color == RED { // case 1
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left { // case 2
					z = z.parent
					T.rightRotate(z)
				}
				z.parent.color = BLACK        // case 3
				z.parent.parent.color = RED   // case 3
				T.leftRotate(z.parent.parent) // case 3
			}
		}
	}
	T.root.color = BLACK
}

func (T *RbTree) RbTreeDelete(z *RbTreeNode) {
	y := z
	var x *RbTreeNode
	yoColor := y.color
	if z.left == nil {
		x = z.right
		T.rbTransplant(z, z.right)

	} else if z.right == nil {
		x = z.left
		T.rbTransplant(z, z.left)

	} else {
		y = RbTreeMinimum(z.right)
		yoColor = y.color
		x = y.right
		if x != nil && y.parent == z {
			x.parent = y
		} else {
			T.rbTransplant(y, y.right)
			y.right = z.right
			if y.right != nil {
				y.right.parent = y
			}

		}

		T.rbTransplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color

	}
	if yoColor == BLACK && x != nil {
		T.rbTreeDeleteFixup(x)
	}
}

func (T *RbTree) rbTransplant(u, v *RbTreeNode) {
	if u.parent == nil {
		T.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}

}

func (T *RbTree) rbTreeDeleteFixup(x *RbTreeNode) {
	for x != T.root && x.color == BLACK {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == RED { // case 1
				w.color = BLACK
				x.parent.color = RED
				T.leftRotate(x.parent)
				w = x.parent.right
			}
			if w.left.color == BLACK && w.right.color == BLACK { // case 2
				w.color = RED
				x = x.parent
			} else {
				if w.right.color == BLACK { // case 3
					w.left.color = BLACK
					w.color = RED
					T.rightRotate(w)
					w = x.parent.right
				}
				w.color = x.parent.color // case 4
				x.parent.color = BLACK   // case 4
				w.right.color = BLACK    // case 4
				T.leftRotate(x.parent)   // case 4
				x = T.root
			}
		} else { // x = x.parent.right
			w := x.parent.left
			if w.color == RED { // case 1
				w.color = BLACK
				x.parent.color = RED
				T.rightRotate(x.parent)
				w = x.parent.left
			}
			if w.left.color == BLACK && w.right.color == BLACK { // case 2
				w.color = RED
				x = x.parent
			} else {
				if w.left.color == BLACK { // case 3
					w.right.color = BLACK
					w.color = RED
					T.leftRotate(w)
					w = x.parent.left
				}
				w.color = x.parent.color // case 4
				x.parent.color = BLACK   // case 4
				w.left.color = BLACK     // case 4
				T.rightRotate(x.parent)  // case 4
				x = T.root
			}
		}
	}
	x.color = BLACK
}
