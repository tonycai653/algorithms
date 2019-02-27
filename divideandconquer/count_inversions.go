// Package main 主要实现了在数组中寻找大小顺序颠倒的数值对, 计算这样的数值对的个数.

// 大小顺序颠倒的对， 比如（5， 2）（7， 5）
// (x, y), x和y都是数值， i和j分别是x和y在数组中的下表位置
// 如果i < j 并且 x > y 那么(x, y)就是数值颠倒的一对

// 程序会读取当前目录下的inter_array.txt文件
// 该文件中每一行有一个数字，一共有100000行。每行的数字不重复， 大小位于[1, 100000]
// 程序会把文件内容读取到一个数组中， 然后计算颠倒对的数目， 并打印出来
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numberArrayFile = "inter_array.txt"

func main() {

	nums, fErr := fileToArray()
	if fErr != nil {
		fmt.Fprintf(os.Stderr, "read numbers in file to slice error: %v\n", fErr)
		os.Exit(1)
	}
	_, numInversions := countInversions(nums)
	fmt.Printf("There are %d inversions in file: %s\n", numInversions, numberArrayFile)

	/*
		a := []uint64{7, 6, 5, 4, 3, 2, 1}

		fmt.Println(countInversions(a))
	*/
}

func fileToArray() ([]uint64, error) {
	f, oErr := os.Open(numberArrayFile)
	if oErr != nil {
		return nil, fmt.Errorf("open file %s error: %v", numberArrayFile, oErr)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	numbers := make([]uint64, 0)
	for scanner.Scan() {
		n, _ := strconv.ParseUint(scanner.Text(), 10, 64)

		numbers = append(numbers, n)
	}
	return numbers, nil
}

func countInversions(nums []uint64) (sortedNums []uint64, numInversions uint64) {
	if len(nums) <= 0 {
		return nil, 0
	}
	if len(nums) == 1 {
		return nums, 0
	}
	mid := len(nums) / 2

	leftSortedNums, leftInversions := countInversions(nums[:mid])
	rightSortedNums, rightInversions := countInversions(nums[mid:])
	sortedNums, mergeInversions := mergeCountInversions(leftSortedNums, rightSortedNums)

	numInversions = leftInversions + rightInversions + mergeInversions

	return
}

func mergeCountInversions(leftSortedNums []uint64, rightSortedNums []uint64) (sortedNums []uint64, numInversions uint64) {
	i, j := 0, 0
	for i < len(leftSortedNums) && j < len(rightSortedNums) {
		if leftSortedNums[i] < rightSortedNums[j] {
			sortedNums = append(sortedNums, leftSortedNums[i])
			i++
		} else if rightSortedNums[j] < leftSortedNums[i] {
			sortedNums = append(sortedNums, rightSortedNums[j])
			j++
			numInversions = numInversions + uint64((len(leftSortedNums) - i))
		} else {
			sortedNums = append(sortedNums, leftSortedNums[i])
			sortedNums = append(sortedNums, rightSortedNums[j])
			i++
			j++
		}
	}
	for i < len(leftSortedNums) {
		sortedNums = append(sortedNums, leftSortedNums[i])
		i++
	}
	for j < len(rightSortedNums) {
		sortedNums = append(sortedNums, rightSortedNums[j])
		j++
	}

	return
}
