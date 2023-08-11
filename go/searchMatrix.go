package main

import (
	"fmt"
)

func main() {
	t := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	b := searchMatrix(t, 3)
	fmt.Println(b)
	fmt.Println()

	t = [][]int{{1}}
	b = searchMatrix(t, 3)
	fmt.Println(b)
	fmt.Println()
}

func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		idx := (high + low) / 2
		mid := nums[idx]
		if mid == target {
			return true
		} else if target < mid {
			high = idx - 1
		} else {
			low = idx + 1
		}
	}

	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	rlow, rhigh := 0, len(matrix)-1

	for rlow <= rhigh {
		idx := (rlow + rhigh) / 2
		v := matrix[idx]
		if target >= v[0] && target <= v[len(v)-1] {
			return search(v, target)
		} else if target < v[0] {
			rhigh = idx - 1
		} else {
			rlow = idx + 1
		}
	}

	return false
}
