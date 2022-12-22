package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%v\n", fizzBuzz(50))
}

// "fizz" if div by 3
// "buzz" if div by 5
// value if not div by 3 or 5
func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 1; i < n+1; i++ {
		if i%3 == 0 {
			ans[i-1] += "Fizz"
		}
		if i%5 == 0 {
			ans[i-1] += "Buzz"
		}
		if ans[i-1] == "" {
			ans[i-1] = strconv.Itoa(i)
		}
	}
	return ans
}
