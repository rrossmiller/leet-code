package main

func maxAreaOfIsland(grid [][]int) int {
	n := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				a := visitBFS(i, j, grid)
				if a > n {
					n = a
				}
			}
		}
	}
	return n
}

func visitBFS(r, c int, grid [][]int) int {
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
