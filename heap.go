package main

import (
	"fmt"
)

// Heap 把数组看成一个二叉树
// 第一个数是树的root
// 假如一个父亲节点的坐标是i, 那么左子树的坐标为2i, 右子树的坐标为2i + 1, (i >= 1)
// 那么如果数组的长度为n, 那么floor(n/2)+1， floor(n/2)+2, ..... 都是叶子节点

// "Max Heap" 的特性： 对于所有树和子树， 它的root的值大于等于起左子树和右子树root的值
// 这样整棵树的root就是最大值

// MaxHeapfy 假设在坐标ind的树的左子树和右子树已经是"Max Heap"
// MaxHeapfy 会对位于坐标ind的树进行操作，使其符合"Max Heap"的特性
func MaxHeapfy(a []int, ind int) {
	l := left(ind)
	r := right(ind)
	largest := ind

	// 我们知道叶子节点的开始位置，所以不对下标进行检查
	if l < len(a) && a[l] > a[largest] {
		largest = l
	}
	if r < len(a) && a[r] > a[largest] {
		largest = r
	}
	if largest != ind {
		swap(a, ind, largest)
		MaxHeapfy(a, largest)
	}
}

// BuildMaxHeap 从无序的数组中构建出"Max Heap"
func BuildMaxHeap(a []int) {
	if len(a) <= 1 {
		return
	}
	// MaxHeap是从下标1开始的
	for i := (len(a) - 1) / 2; i >= 1; i-- {
		MaxHeapfy(a, i)
	}
}

func HeapSort(a []int) {
	BuildMaxHeap(a)
	for i := len(a) - 1; i >= 2; i-- {
		swap(a, i, 1)
		a = a[:len(a)-1]
		MaxHeapfy(a, 1)
	}
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func swap(a []int, i, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

func main() {
	// 第一个位置为0， 占位，不算在数据内
	a := []int{0, 2, 8, 3, 5, 2, 6, 4, 1, 12}

	HeapSort(a)
	fmt.Println(a)
}
