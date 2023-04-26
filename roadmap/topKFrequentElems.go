package main

import (
	"fmt"
	"math"
	"sort"
)

type Val struct {
	value int
	cnt   int
}

// Given an integer array nums and an integer k,
// return the k most frequent elements.
func topKFrequent(nums []int, k int) []int {
	ans := make([]Val, 0)
	sort.Ints(nums)
	fmt.Println(nums)

	curr := math.MinInt
	count := 0
	for _, v := range nums {
		if v != curr {
			if curr != math.MinInt {
				ans = append(ans, Val{value: curr, cnt: count})
			}
			curr = v
			count = 1
		} else {
			count++
		}
	}
	ans = append(ans, Val{value: curr, cnt: count})

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].cnt > ans[j].cnt
	})

	fmt.Println(ans)
	rtn := make([]int, k)
	for i := 0; i < k; i++ {
		rtn[i] = ans[i].value
	}
	return rtn
}

func main() {
	c := []int{1, 1, 2, 1, 2, 3}
	k := 2
	x := topKFrequent(c, k)
	fmt.Printf("ans %v", x)
}
