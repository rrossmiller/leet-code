package main

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	res := nums[0]

	for l <= r {
		if nums[l] < nums[r] {
			return min(res, nums[l])
		}

		m := (l + r) / 2
		res = min(res, nums[m])

		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}

	}

	return res
}
