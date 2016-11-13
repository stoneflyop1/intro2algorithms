# 排序算法
主要参考了《算法导论》，其中[sorts.go](sorts.go)为测试代码

所有线性表元素个数都假定为n

## 选择排序(SelectionSort)

代码实现见[SelectionSort.go](sorts/SelectionSort.go)，参考《算法设计与分析基础》第3.1节

1. 从待排序线性表中根据两两比较找到最大的元素，把第一个元素与最小的元素做交换；
2. 因为第一个元素已经是最小元素，位置确定，所以从线性表的第二个元素开始找其中最小的元素(第二小元素)，再把第二个元素跟找到的最小元素交换；
3. 按照欧第1、2步骤的方法依次找到第m个最小元素跟第m个位置的元素进行交换，直到m=n-1

伪代码如下(使用数组实现)：

```
SelectionSort(A[0...n-1])
    //输入: 数组 A[0...n-1]，其中元素是可排序的(即：可比较的)
    //输出：非降序的数组，修改了原数组
    for i = 0 to n-2 do
        min = i
        for j = i+1 to n-1 do
            if A[j] < A[min] // 此处没有等号，所以程序是稳定的
                min = j
        A[i],A[min] = A[min],A[i]
```

从伪代码中可以看到，元素之间的比较为平方级效率 n<sup>2</sup>，而交换元素则是线性效率`n`

## 冒泡排序(BubbleSort)

代码实现见[BubbleSort.go](sorts/BubbleSort.go)，参考《算法设计与分析基础》第3.1节

1. 从第1个元素开始，与第2个元素进行比较，若大于第2个元素，则交换；若第2个元素比第3个元素大，则交换；一直到n-1和n；最大值交换到n的位置
2. 与第一步类似，从第1个元素到第n-1个元素；一直到从第1个元素到第2个元素截止

伪代码如下(使用数组实现)：

```
BubbleSort(A[0...n-1])
    for i = 0 to n-2 do
        for j = 0 to n-2-i do
            if A[j+1] < A[j]
                A[j+1],A[j] = A[j],A[j+1]
```

比较的效率为平方级 n<sup>2</sup>，交换元素的效率也是平方级 n<sup>2</sup>

## 插入排序(InsertionSort)

代码实现[InsertionSort.go](sorts/InsertionSort.go)，参考《算法设计与分析基础》第4.1节

与打扑克类比：
* 插入排序就像打扑克发牌时，把牌按照大小顺序一个一个的插入到手中。
* 手里没牌时，直接放到手中；
* 有一个牌时，先比较发到的牌与手中的牌的大小，若手中牌大，则插入左侧；若手中牌小，则插入右侧。
* 若手中牌较多时，则可能会按照大小顺序插入中间。

伪代码如下(使用数组实现)：

```
InsertionSort(A[0...n-1])
    for i = 1 to n-1 do
        v = A[i]
        j = i-1
        while j >= 0 && A[j] > v do
            A[j+1] = A[j] //位置小于i的元素已排好序，因此只需依次移位
            j = j-1
        A[j+1] = v
```

对于接近排好序的线性表，插入排序拥有很好的效率

## 归并排序(MergeSort)

代码实现[MergeSort.go](sorts/MergeSort.go)，递归版本参考《算法导论》第2.3.1节

* 分(Divide)：把n个元素的线性表序列划分成两个子序列，每个子序列含有n/2个元素
* 治(Conquer)：使用归并方法对两个子序列进行递归排序
* 和(Combine)：把排好序的两个子序列合并为一个排好序的序列
在不断递归的最后是n个长度为1的子序列，然后再依次向上排序合并。

伪代码如下(参考《算法设计与分析导论》第5.1节)：

```
MergeSort(A[0...n-1])
    if n > 1
        copy A[0...n/2-1] to B[0...n/2-1]
        copy A[n/2-1...n-1] to C[0...n/2-1]
        MergeSort(B[0...n/2-1])
        MergeSort(C[0...n/2-1])
        Merge(B,C,A)

Merge(B[0...p-1], C[0...q-1], A[0...p-q+1])
    //输入：已排好序的数组B[0...p-1]和C[0...q-1]
    //输出：B和C合并为排好序的数组A[0...p+q-1]
    i = 0; j = 0; k = 0
    while i < p && j < q  do
        if B[i] <= C[j]
            A[k] = B[j]; i = i + 1
        else
            A[k] = C[j]; j = j + 1
        k = k + 1
    if i == p
        copy C[j...q-1] to A[k...p+q-1]
    else
        copy B[i...p-1] to A[k...p+q-1]

```

## (最大)堆排序(HeapSort)

主要参考《算法导论》第6章。

一个二叉堆(binary heap)可以看作一个以数组(array)方式实现的几近完全的二叉树(binary tree)。
最大堆特性：父节点值不小于子节点的值，即：根元素最大。

### 父节点和左右子节点的编号关系

若编号以1开始，则：
* i的父节点编号不大于i/2
* i的左子节点编号为2*i
* i的右子节点编号为2*i+1

注：编号以0开始的请参考代码实现

主程序包含三部分：

* MaxHeapify —— 用来保持最大堆特性，对数时间
* BuildMaxHeap —— 从未排序数组中构造一个最大堆，线性时间
* HeapSort —— 对数组进行在位排序，n*log n时间

程序从BuildMaxHeap开始先构造一个最大堆，这样第一个元素为最大。把第一个元素跟最后一个元素交换，通过MaxHeapify使第一个元素到第n-1个元素保持最大堆特性，依次循环；按照此种方式进行到第二个元素，就得到了一个非降序的数组。

```
HeapSort (A)
    BuildMaxHeap(A)
    for i = A.length downto 2
        A[1],A[i] = A[i],A[1]
        A.heapSize = A.heapSize-1
        MaxHeapify(A, 1) // 1st is always the largest one
```

## 快速排序(QuickSort)

参考《算法导论》第7章和《算法设计与分析基础》第5.2节。

快速排序也是一种分治算法。与归并排序类似，也需要对数组进行划分，不过与其不同的是，不是根据位置，而是根据元素值划分。做过一次划分之后，选定的划分元素则已经在了最终的位置。与归并排序另一个不同的地方时，快速排序不需要合并，所有工作都在划分程序中完成。

假设对数组的A[p...r]元素进行排序，程序名为QuickSort：
* 分(Divide)：选定一个q(p<=q<=r)，把数组划分为两个子数组Left=A[p...q-1], Right=A[q+1...r]，使得Left中的所有元素都不大于A[q]，Right中所有元素都不小于A[q]
* 治(Conquer)：对Left和Right递归调用QuickSort
* 合(Combine)：无需合并

```
QuickSort(A, p, r)
    if p < r
        q = partition(A, p, r)
        QuickSort(A, p, q-1)
        QuickSort(A, q+1, r)
```

注：《算法设计与分析基础》第5.2节给出了一个更高效的划分算法(Hoare)；划分算法还可以用在选择问题(Selection Problem)中(选择第k小的元素问题)，代码实现见代码文件中的Selection4K。划分程序还需要注意循环变量的边界问题(数组下标越界问题)。
