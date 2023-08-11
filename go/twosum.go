package main

import (
	"fmt"
	"sort"
)

func main() {
	// 	Input: nums = [2,7,11,15], target = 9
	// Output: [0,1]
	nums := []int{2, 7, 11, 15}
	tgt := 9
	fmt.Println(twoSum(nums, tgt))
}

func twoSum(nums []int, tgt int) []int {
	nums = sort.IntSlice(nums)
	for i, v := range nums {
		rem := tgt - v
		j := i + 1
		for _, x := range nums[i+1:] {
			if x == rem {
				return []int{i, j}
			}
			j++
		}
	}
	return nil
}
