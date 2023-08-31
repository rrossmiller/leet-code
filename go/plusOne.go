package main

import (
	"fmt"
)

func main() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits))
	digits = []int{9}
	fmt.Println(plusOne(digits))

	digits = []int{9, 9}
	fmt.Println(plusOne(digits), 100)

	digits = []int{9, 9, 9}
	fmt.Println(plusOne(digits), 1000)

	digits = []int{9, 0, 9}
	fmt.Println(plusOne(digits), 910)

	digits = []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	fmt.Println(plusOne(digits))

}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			break
		}

	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
