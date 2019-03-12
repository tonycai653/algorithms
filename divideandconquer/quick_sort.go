package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const numberArrayFile = "quick_sort.txt"

func fileToArray() ([]int64, error) {
	f, oErr := os.Open(numberArrayFile)
	if oErr != nil {
		return nil, fmt.Errorf("open file %s error: %v", numberArrayFile, oErr)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	numbers := make([]int64, 0)
	for scanner.Scan() {
		n, _ := strconv.ParseInt(scanner.Text(), 10, 0)

		numbers = append(numbers, n)
	}
	return numbers, nil
}

func main() {
	a, fErr := fileToArray()
	if fErr != nil {
		fmt.Fprintf(os.Stderr, "read numbers to array for quick sort: %v\n", fErr)
		os.Exit(1)
	}

	comps := QuickSort(a)
	fmt.Println(comps)

}

func QuickSort(n []int64) uint64 {
	if len(n) <= 1 {
		return 0
	}
	pivot, comparisons := partition(n)
	leftComparisons := QuickSort(n[:pivot])
	rightComparisons := QuickSort(n[pivot+1:])

	return comparisons + leftComparisons + rightComparisons
}

func partition(n []int64) (pivotIndex int64, numberComparison uint64) {
	finalIndex := len(n) - 1
	pivot := finalIndex

	if pivot != finalIndex {
		swap(n, finalIndex, pivot)
	}
	pivotNumber := n[finalIndex]
	i := 0
	for j := 0; j < len(n); j++ {
		if n[j] < pivotNumber {
			numberComparison++
			swap(n, i, j)
			i++
		}
	}
	swap(n, i, finalIndex)
	pivotIndex = int64(i)
	return
}

func swap(n []int64, i, j int) {
	t := n[i]
	n[i] = n[j]
	n[j] = t
}
