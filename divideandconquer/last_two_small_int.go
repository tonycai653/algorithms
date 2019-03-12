package main

import (
	"fmt"
	"math"
)

// Sort找出一个数组中的最小的两个数
func Smallest(a []int) (x, y int) {
	if len(a) == 1 {
		return a[0], math.MaxInt64
	}
	if len(a) == 2 {
		return a[0], a[1]
	}
	mid := len(a) / 2
	leftArray := a[:mid]
	rightArray := a[mid:]

	leftSmallX, leftSmallY := Smallest(leftArray)
	rightSmallX, rightSmallY := Smallest(rightArray)

	x, left := smallestOfFour(&leftSmallX, &leftSmallY, &rightSmallX, &rightSmallY)
	y = smallestOfThree(left[0], left[1], left[2])

	return
}

func smallestOfThree(a, b, c int) int {
	s := a
	if a < b {
		s = a
	} else {
		s = b
	}
	if s < c {
		return s
	}
	return c
}

func smallestOfFour(a, b, c, d *int) (s int, left [3]int) {

	lx, ly := smaller(a, b)
	rx, ry := smaller(c, d)
	t, _ := smaller(lx, rx)
	s = *t

	left[0] = *ly
	left[1] = *ry
	if t == lx {
		left[2] = *rx
	} else {
		left[2] = *lx
	}
	return
}

func smaller(a *int, b *int) (*int, *int) {
	if *a < *b {
		return a, b
	}
	return b, a
}

func main() {
	a := []int{20, 10, 3, 4, -1, 0}
	fmt.Println(Smallest(a))
}
