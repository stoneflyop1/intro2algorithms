# 一般数据结构

注：主要参考了《算法导论》，其中[datastructs.go](datastructs.go)为测试代码。

## 栈(stack)和队列(queue)


## 树节点的表示
### 二叉树(binary tree)

```
    type BTreeNode struct {
        key    interface{}
        parent *BTreeNode
        left   *BTreeNode
        right  *BTreeNode
    }
```

二叉树遍历有三种形式：
* 前序遍历(preorder traversal 或者 preorder tree walk): 根->左子树->右子树
* 中序遍历(inorder traversal 或者 inorder tree walk)：左子树->根->右子树
* 后序遍历(postorder traversal 或者 postorder tree walk)：左子树->右子树->根

三种遍历形式何时使用，见：http://stackoverflow.com/questions/9456937/when-to-use-preorder-postorder-and-inorder-binary-search-tree-traversal-strate

```
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
```

可以很容易的用递归方式实现上述三种遍历方式，如何在不使用栈和队列等其他数据结构来实现非递归方式呢？

### 一般树
可以用类似二叉树的节点形式表示一颗一般树的节点。node.leftChild表示node的最左子节点，node.rightSibling表示紧挨着最左节点的子节点，而node.rightSibling.rightSibling按从左到右的顺序依次指向其余的子节点

```
    type TreeNode struct {
        key          interface{}
        parent       *TreeNode
        leftChild    *TreeNode
        rightSibling *TreeNode
    }
```

## 二叉搜索树(binary search tree)
任何一个节点n，其左子树的节点值都不大于节点n的值，其右子树的节点值都不小于节点n的值



