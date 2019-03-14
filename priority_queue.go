package main

import (
	"fmt"
)

type MaxPriorityQueue struct {
	nums []int
}

func New(n []int) *MaxPriorityQueue {
	q := &MaxPriorityQueue{
		nums: n,
	}
	q.buildMaxHeap()
	return q
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

func (q *MaxPriorityQueue) maxHeapfy(ind int) {
	l := left(ind)
	r := right(ind)
	largest := ind

	// 我们知道叶子节点的开始位置，所以不对下标进行检查
	if l < len(q.nums) && q.nums[l] > q.nums[largest] {
		largest = l
	}
	if r < len(q.nums) && q.nums[r] > q.nums[largest] {
		largest = r
	}
	if largest != ind {
		swap(q.nums, ind, largest)
		q.maxHeapfy(largest)
	}
}

// buildMaxHeap 从无序的数组中构建出"Max Heap"
func (q *MaxPriorityQueue) buildMaxHeap() {
	if len(q.nums) <= 1 {
		return
	}
	// MaxHeap是从下标1开始的
	for i := (len(q.nums) - 1) / 2; i >= 1; i-- {
		q.maxHeapfy(i)
	}
}

func (q *MaxPriorityQueue) Maximum() int {
	return q.nums[1]
}

func (q *MaxPriorityQueue) ExtractMaximum() int {
	// 第一个位置是占位
	if len(q.nums) < 2 {
		panic("heap underflow")
	}
	max := q.nums[1]
	q.nums[1] = q.nums[len(q.nums)-1]
	q.nums = q.nums[:len(q.nums)-1]
	q.maxHeapfy(1)
	return max
}

func (q *MaxPriorityQueue) IncreaseKey(i, key int) {
	if i <= 1 || i > len(q.nums)-1 {
		panic("index out of range")
	}
	if key < q.nums[i] {
		panic("new key is smaller than current key")
	}

	q.nums[i] = key
	for i > 1 && q.nums[i/2] < q.nums[i] {
		swap(q.nums, i, i/2)
		i = i / 2
	}
}

func (q *MaxPriorityQueue) Insert(key int) {
	q.nums = append(q.nums, key-1)
	q.IncreaseKey(len(q.nums)-1, key)
}

func main() {
	maxQueue := New([]int{1, 2, 3, 4, 5})
	fmt.Println(maxQueue.Maximum())
	fmt.Println(maxQueue.ExtractMaximum())
	fmt.Println(maxQueue.Maximum())
	maxQueue.Insert(20)
	fmt.Println(maxQueue.Maximum())
	// Output:
	// 5
	// 5
	// 4
	// 20
}
