package main

import "fmt"

func main() {
	a := search([]int{1, 2, 3, 4}, 3)
	fmt.Println(a)
	a = search([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Println(a)
	a = search([]int{-1, 0, 3, 5, 9, 12}, 12)
	fmt.Println(a)

	fmt.Println()
	a = search([]int{-1, 0, 3, 5, 9, 12}, 11)
	fmt.Println(a == -1)
	fmt.Println()
	a = search([]int{-1, 2, 5}, 5)
	fmt.Println(a)
}

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		idx := (high + low) / 2
		mid := nums[idx]
		if mid == target {
			return idx
		} else if target < mid {
			high = idx - 1
		} else {
			low = idx + 1
		}
	}

	return -1
}
