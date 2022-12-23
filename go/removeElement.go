package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//https://leetcode.com/problems/remove-element/
	nums := []int{3, 2, 2, 3}
	val := 3
	ans := removeElement(nums, val)
	fmt.Println(ans)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	ans = removeElement(nums, val)
	fmt.Println(ans)

}

func removeElement(nums []int, val int) (k int) {
	// remove val from nums
	// return len(nums) - (num times val was in nums)
	k = 0
	for i, _ := range nums {
		if nums[i] == val {
			k++
			nums[i] = math.MaxInt64
		}
	}
	sort.Ints(nums)

	k = len(nums) - k

	return
}
