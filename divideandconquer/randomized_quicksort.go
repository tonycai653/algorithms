package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randIntBetween(lower, higher int) int {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	return r.Intn(higher-lower) + lower
}

func randomizedQuickSort(n []int) {
	if len(n) <= 1 {
		return
	}
	pivot := partition(n)
	randomizedQuickSort(n[:pivot])
	randomizedQuickSort(n[pivot+1:])
}

func partition(n []int) int {
	finalIndex := len(n) - 1
	pivot := randIntBetween(0, finalIndex)

	if pivot != finalIndex {
		swap(n, finalIndex, pivot)
	}
	pivotNumber := n[finalIndex]
	i := 0
	for j := 0; j < len(n); j++ {
		if n[j] < pivotNumber {
			swap(n, i, j)
			i++
		}
	}
	swap(n, i, finalIndex)
	return i
}

func swap(n []int, i, j int) {
	t := n[i]
	n[i] = n[j]
	n[j] = t
}

func main() {

	a := []int{2, 8, 3, 20, 18, 10, 1}

	randomizedQuickSort(a)

	fmt.Println(a)
	// Output:
	// [1, 2, 3, 8, 10, 18, 20]
}
