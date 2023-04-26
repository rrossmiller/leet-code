package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	c := 1
	for i := 0; i < len(nums); i++ {
		res[i] = c
		c = c * nums[i]
	}

	c = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * c
		c = c * nums[i]
	}

	return res
}

func main() {
	in := []int{1, 2, 3, 4}
	fmt.Println(in)
	x := productExceptSelf(in)
	fmt.Println(x)
	fmt.Println()

	in = []int{-1, 1, 0, -3, 3}
	fmt.Println(in)
	x = productExceptSelf(in)
	fmt.Println(x)
}
