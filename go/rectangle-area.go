package main

import "fmt"

// Given the coordinates of two rectilinear rectangles in a 2D plane, return the total area covered by the two rectangles.
// The first rectangle is defined by its bottom-left corner (ax1, ay1) and its top-right corner (ax2, ay2).
// The second rectangle is defined by its bottom-left corner (bx1, by1) and its top-right corner (bx2, by2).

func main() {
	ax1 := -3
	ay1 := 0
	// ax2 := 3
	// ay2 := 4
	// bx1 := 0
	// by1 := -1
	// bx2 := 9
	// by2 := 2

	a(ax1, ax1, ay1)
}

func a(x ...int) {
	for _, v := range x {
		fmt.Println(v)
	}
}
