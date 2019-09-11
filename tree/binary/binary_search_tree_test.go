package binary_test

import (
	"sort"
	"testing"

	"github.com/tonycai653/algorithms/tree/binary"
)

func TestBinarySearchTree(t *testing.T) {
	var tree *binary.SearchTree

	t.Run("recursive insert", func(t *testing.T) {
		nums := []int{30, 10, 20, 18, 2, 6, 9, 11, 66}

		tree = binary.NewTree(nums)

		actualNums := tree.InorderWalk()
		if len(actualNums) != len(nums) {
			t.Fatalf("Expected equal length of numbers, Expected: %d, Got: %d\n", len(nums), len(actualNums))
		}
		sort.Ints(nums)
		for i := 0; i < len(nums); i++ {
			if nums[i] != actualNums[i] {
				t.Fatalf("Expected nums and actualNums equal, but not\n")
			}
		}
	})

	t.Run("iterative minimum", func(t *testing.T) {
		if min := tree.IterativeMinimum(); min.GetKey() != 2 {
			t.Fatalf("Expected mimium number: %d, Got: %d\n", 2, min.GetKey())
		}
	})

	t.Run("recursive minimum", func(t *testing.T) {
		if min := tree.RecursiveMinimum(); min.GetKey() != 2 {
			t.Fatalf("Expected mimium number: %d, Got: %d\n", 2, min.GetKey())
		}
	})

	t.Run("iterative maximum", func(t *testing.T) {
		if max := tree.IterativeMaximum(); max.GetKey() != 66 {
			t.Fatalf("Expected maximum number: %d, Got: %d\n", 66, max.GetKey())
		}
	})

	t.Run("recursive maximum", func(t *testing.T) {
		if max := tree.RecursiveMaximum(); max.GetKey() != 66 {
			t.Fatalf("Expected maximum number: %d, Got: %d\n", 66, max.GetKey())
		}
	})

	t.Run("successor", func(t *testing.T) {
		if suc := tree.Successor(); suc.GetKey() != 66 {
			t.Fatalf("Expected successor number: %d, Got: %d\n", 66, suc.GetKey())
		}
	})

	t.Run("predecessor", func(t *testing.T) {
		if suc := tree.Predecessor(); suc.GetKey() != 20 {
			t.Fatalf("Expected predecessor number: %d, Got: %d\n", 20, suc.GetKey())
		}
	})

	t.Run("search tree", func(t *testing.T) {
		if node := tree.Find(66); node == nil {
			t.Fatalf("Expected 66 found, but not\n")
		}
		if node := tree.Find(-1); node != nil {
			t.Fatalf("Expected -1 not found, but found")
		}
	})

	t.Run("delete", func(t *testing.T) {
		tree.Delete(30)
		if node := tree.Find(30); node != nil {
			t.Fatalf("Expected 30 not found after deletion")
		}
		if root := tree.Root(); root.GetKey() != 66 {
			t.Fatalf("Expected 66 as root key, but Got: %d", root.GetKey())
		}
	})
}
