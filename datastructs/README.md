# 一般数据结构

注：主要参考了《算法导论》，其中[datastructs.go](datastructs.go)为测试代码。

## 栈和队列

注：代码的数组实现中考虑了数组下溢和上溢两种异常，而不像伪代码那样：只有栈考虑下溢异常。

### 栈(stack)

栈是一个先进后出(LIFO)的数据结构，一般仅有入栈和出栈两个操作和判断栈是否为空操作。栈为空时，无法执行出栈操作。可以使用数组或者链表实现。
下面为以数组实现的伪代码：

```
StackIsEmpty(S)
    if S.top == 0
        return true
    return false

Push(S, x)
    S.top++
    S[S.top] = x

Pop(S)
    if StackIsEmpty(S)
        error "underflow"
    S.top--
    return S[S.top+1]

```

### 队列(queue)

队列是一个先进先出(FIFO)的数据结构，一般有入队和出队两个操作。可以使用数组或者链表实现。
数组实现的伪代码：

```
Enqueue(Q, x)
    Q[Q.tail] = x
    if Q.tail == Q.length
        Q.tail = 1
    else
        Q.tail++

Dequeue(Q)
    x = Q[Q.head]
    if Q.head == Q.length
        Q.head = 1
    else
        Q.head++
    return x
```

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


## 红黑二叉树

与《算法导论》中伪码稍有不同，go代码实现未使用哨兵NIL

### 节点特性

* 每个节点或者是红的，或者是黑的
* 树的根是黑的
* 每个扩展的空叶子(NIL)都是黑的
* 若节点是红的，则它的两个子节点都是黑的
* 对于所有节点，所有到其子孙的简单路径都含有同样个数的黑节点（black-height）

### 树特性

* 具有n个节点的红黑树的高度至多为`2Xlog(n+1)` (引理13.1)



