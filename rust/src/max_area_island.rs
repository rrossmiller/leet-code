use std::collections::VecDeque;

struct Solution;

impl Solution {
    pub fn max_area_of_island(grid: Vec<Vec<i32>>) -> i32 {
        let mut n = 0;
        let mut visited = Self::get_visited(grid.len(), grid[0].len());

        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                if grid[i][j] == 1 && !visited[i][j] {
                    let a = Self::visit_bfs(i, j, &grid, &mut visited);
                    if a > n {
                        n = a;
                    }
                }
            }
        }

        n
    }

    fn visit_bfs(r: usize, c: usize, grid: &Vec<Vec<i32>>, visited: &mut Vec<Vec<bool>>) -> i32 {
        let mut a = 1;
        let mut q = VecDeque::new();
        q.push_front(vec![r, c]);
        visited[r][c] = true;

        let dirs = vec![vec![1, 0], vec![-1, 0], vec![0, 1], vec![0, -1]];

        while q.len() > 0 {
            let x = q.pop_front().unwrap();
            let r = x[0];
            let c = x[1];

            for d in dirs.iter() {
                let x = r as i32 + d[0];
                let y = c as i32 + d[1];

                if x >= 0
                    && x < grid.len() as i32
                    && y >= 0
                    && y < grid[0].len() as i32
                    && grid[x as usize][y as usize] == 1
                    && !visited[x as usize][y as usize]
                {
                    a += 1;
                    q.push_front(vec![x as usize, y as usize]);
                    visited[x as usize][y as usize] = true;
                }
            }
        }
        a
    }

    fn get_visited(r: usize, c: usize) -> Vec<Vec<bool>> {
        let mut v = vec![];
        for _ in 0..r {
            v.push(vec![false; c])
        }

        v
    }
}
