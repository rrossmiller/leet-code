package main

import "fmt"

func main() {
	twoSum(nil, 0)
}

func twoSum(nums []int, tgt int) []int {
	prev := make([]int, len(nums))
	fmt.Println(prev)
	for i, x := range nums {
		rem := tgt - x
		if contains(prev, rem) {
			return []int{i, prev[rem]}
		}

		prev[x] = i
	}
	return nil
}

func contains(prev []int, idx int) bool {
	return prev[idx] > 0
}
