package main

import (
	"fmt"
)

// func main() {
// 	var a, b string
// 	// Input:
// 	a = "11"
// 	b = "1"
// 	fmt.Println("a:", a)
// 	fmt.Println("b:", b)
// 	fmt.Println(addBinary(a, b))
// 	fmt.Println("100")
//
// 	fmt.Println("************")
// 	a = "1010"
// 	b = "1011"
// 	fmt.Println("a:", a)
// 	fmt.Println("b:", b)
// 	fmt.Println(addBinary(a, b))
// 	fmt.Println("10101")
//
// 	fmt.Println("************")
// 	a = "1111"
// 	b = "1111"
// 	fmt.Println("a:", a)
// 	fmt.Println("b:", b)
// 	fmt.Println(addBinary(a, b))
// 	fmt.Println("11110")
// }

func addBinary(a string, b string) string {
	var ans, A, B string
	lenA := len(a) - 1
	lenB := len(b) - 1

	carry := false

	for lenA >= 0 || lenB >= 0 {
		place := "0"
		A = ""
		B = ""
		if carry {
			place = "1"
			carry = false
		}
		if lenA >= 0 {
			A = string(a[lenA])
			lenA--
			place = check(place, A, &carry)
		}

		if lenB >= 0 {
			B = string(b[lenB])
			lenB--
			place = check(place, B, &carry)
		}

		ans = place + ans
		fmt.Println("a:", ans)
	}
	if carry {
		ans = "1" + ans
	}
	return ans
}

func check(place, s string, carry *bool) string {
	out := place
	if place == "1" && s == "1" {
		out = "0"
		*carry = true
	} else if s == "1" {
		out = "1"
		if !*carry {
			*carry = false
		}
	}
	fmt.Printf("%v %v -> p:%v\n", place, s, out)
	return out
}
