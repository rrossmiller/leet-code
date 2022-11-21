package main

import (
	"fmt"
	"os"
	"strconv"
)

// Given the coordinates of two rectilinear rectangles in a 2D plane, return the total area covered by the two rectangles.
// The first rectangle is defined by its bottom-left corner (ax1, ay1) and its top-right corner (ax2, ay2).
// The second rectangle is defined by its bottom-left corner (bx1, by1) and its top-right corner (bx2, by2).

func main() {
	ax1, _ := strconv.Atoi(os.Args[1])
	ay1, _ := strconv.Atoi(os.Args[2])
	ax2, _ := strconv.Atoi(os.Args[3])
	ay2, _ := strconv.Atoi(os.Args[4])
	bx1, _ := strconv.Atoi(os.Args[5])
	by1, _ := strconv.Atoi(os.Args[6])
	bx2, _ := strconv.Atoi(os.Args[7])
	by2, _ := strconv.Atoi(os.Args[8])
	// fmt.Println(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
	// fmt.Println()
	res := computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
	fmt.Println(res)
}

type Point struct {
	x int
	y int
}

func (p Point) invalid(v *Point) bool {
	if v == nil {
		return false
	}
	return p.x == v.x || p.y == v.y
}

type Rectangle struct {
	TL Point
	TR Point
	BL Point
	BR Point
}

func (r *Rectangle) area() int {
	l := r.TL.x - r.TR.x
	if l < 0 {
		l *= -1
	}
	h := r.TL.y - r.BL.y
	if h < 0 {
		h *= -1
	}
	return l * h
}

// Edges defined clockwise from the top
// Lines defined L-R, T-B
func (r *Rectangle) getEdges() [][]Point {
	top := []Point{r.TL, r.TR}
	right := []Point{r.TR, r.BR}
	bottom := []Point{r.BL, r.BR}
	left := []Point{r.TL, r.BL}

	return [][]Point{top, right, bottom, left}
}

func computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) int {
	fmt.Println(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 )
	r1 := getRect(ax1, ay1, ax2, ay2)
	r2 := getRect(bx1, by1, bx2, by2)
	// fmt.Println(r1)
	// fmt.Println(r2)
	// fmt.Println()

	ri := findIntersection(r1, r2)
	if ri != nil {
		return r1.area() + r2.area() - ri.area()
	}
	return r1.area() + r2.area()
}

func getRect(x1, y1, x2, y2 int) Rectangle {
	bl := Point{x1, y1}
	tr := Point{x2, y2}
	tl := Point{x1, y2}
	br := Point{x2, y1}

	return Rectangle{tl, tr, bl, br}
}

func findIntersection(r1, r2 Rectangle) *Rectangle {
	// check if any of the recangles' edges intersect
	r1Edges := r1.getEdges()
	r2Edges := r2.getEdges()

	//todo check if a part of a rectangle is inside the other
	// r1 right interects r2 top or bottom
	p1 := checkIntersect(r1Edges[1], r2Edges[0], "right1", "top2")
	if p1 == nil {
		p1 = checkIntersect(r1Edges[1], r2Edges[2], "right1", "bottom2")
	}
	// r1 left interects r2 top or bottom
	if p1 == nil {
		p1 = checkIntersect(r1Edges[3], r2Edges[0], "left1", "top2")
	}
	if p1 == nil {
		p1 = checkIntersect(r1Edges[3], r2Edges[2], "left1", "bottom2")
	}
	// r1 top interects r2 right or left
	p2 := checkIntersect(r2Edges[1], r1Edges[0], "right2", "top1")
	if p2 == nil || p1.invalid(p2) {
		p2 = checkIntersect(r2Edges[3], r1Edges[0], "left2", "top1")
	}

	// r1 bottom interects r2 right or left
	if p2 == nil || p1.invalid(p2) {
		p2 = checkIntersect(r2Edges[1], r1Edges[2], "right2", "bottom1")
	}
	if p2 == nil || p1.invalid(p2) {
		p2 = checkIntersect(r2Edges[3], r1Edges[2], "left2", "bottom1")
	}

	fmt.Println("results")
	fmt.Println(p1)
	fmt.Println(p2)
	if p1 != nil && p2 != nil {
		rtn := getRect(p1.x, p1.y, p2.x, p2.y) // doesn't follow top right bottomw left notation
		fmt.Println(rtn)
		return &rtn
	}
	return nil
}

func checkIntersect(vert, horiz []Point, vertName, horizName string) *Point {
	// vertical x is between the endpoints of the horizontal
	// AND horizontal y is between the endpoints of the vertical
	if (vert[0].x >= horiz[0].x && vert[0].x <= horiz[1].x) && (horiz[0].y <= vert[0].y && horiz[0].y >= vert[1].y) {
		// fmt.Println(vertName, horizName)
		// fmt.Println(vert, horiz)

		return &Point{vert[0].x, horiz[0].y}
	}
	return nil
}
