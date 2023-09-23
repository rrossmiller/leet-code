package main

import "fmt"

func main() {
	image := [][]int{[]int{1, 1, 1}, []int{1, 1, 0}, []int{1, 0, 1}}
	sr := 1
	sc := 1
	color := 2
	ans := [][]int{[]int{2, 2, 2}, []int{2, 2, 0}, []int{2, 0, 1}}
	output := floodFill(image, sr, sc, color)
	checkAns(output, ans)

	image = [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}
	sr = 1
	sc = 1
	color = 0
	output = floodFill(image, sr, sc, color)
	checkAns(output, image)

	// [[0,0,0],[0,1,0]]
}

func checkAns(output, ans [][]int) {
	fmt.Println(output)
	fmt.Println(ans)
	ok := true
	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(ans[0]); j++ {
			if ans[i][j] != output[i][j] {
				fmt.Println(ans[i][j], output[i][j])
				ok = false
				break
			}
		}
	}

	fmt.Println(ok)
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	origColor:=image[sr][sc] 
	image[sr][sr]=color

	// BFS
	q := [][2]int{{sr, sc}}
	dirs := [4][2]int{
		{1, 0},  // up
		{-1, 0}, //down
		{0, 1},  //right
		{0, -1}, //left
	}

	// while the queue is not empty
	for len(q) > 0 {
		//pop off the queue
		r, c := q[0][0], q[0][1]
		q = q[1:]
		// for every direction
		for _, d := range dirs {
			// get the grid coordinates to check
			x := r + d[0] 
			y := c + d[1]

			// check that the coordinates are within bounds
			if x >= 0 && x < len(image) && y >= 0 && y < len(image[x]) && image[x][y] == origColor {
				q = append(q, [2]int{x, y})
				image[x][y] = color
			}
		}
	}
	return image

}

func BFS(r, c int, grid [][]int) int {
	a := 1
	q := [][2]int{{r, c}}
	grid[r][c] = '.'
	dirs := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for len(q) > 0 {
		r, c := q[0][0], q[0][1]
		q = q[1:]
		for _, d := range dirs {
			x := r + d[0]
			y := c + d[1]

			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && grid[x][y] == 1 {
				q = append(q, [2]int{x, y})
				grid[x][y] = -1
				a += 1
			}
		}
	}
	return a
}
