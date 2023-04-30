package main

import (
	"fmt"
)

func main() {
	p := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(p))
	fmt.Println()

	p = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(p))
	fmt.Println()
}

func maxProfit(prices []int) int {
	rtn := 0
	min := prices[0]
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > rtn {
			rtn = prices[i] - min
		}

	}
	return rtn
}
