package twentyDays

import (
	"math"
)

/* 542. 01 矩阵
https://leetcode-cn.com/problems/01-matrix/
*/
func updateMatrix(mat [][]int) [][]int {
	ret := make([][]int, len(mat))
	queue := make([][2]int, 0)
	for i := 0; i < len(mat); i++ {
		ret[i] = make([]int, len(mat[0]))
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
				ret[i][j] = 0
			} else {
				ret[i][j] = math.MaxInt32
			}
		}
	}

	dist := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			sr := tmp[0] + dist[i][0]
			sc := tmp[1] + dist[i][1]
			if sc < 0 || sr < 0 || sr > len(mat)-1 || sc > len(mat[0])-1 {
				continue
			}

			if ret[sr][sc] > ret[tmp[0]][tmp[1]]+1 {
				ret[sr][sc] = ret[tmp[0]][tmp[1]] + 1
				queue = append(queue, [2]int{sr, sc})
			}
		}
	}
	return ret
}

/* 994. 腐烂的橘子
https://leetcode-cn.com/problems/rotting-oranges/submissions/
 */
func orangesRotting(grid [][]int) int {
	queue := make([][2]int, 0)
	scan := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				scan++
			}
		}
	}

	if scan == 0 {
		return 0
	}

	ret := 0
	dist := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for len(queue) > 0 {
		ret++
		tmpSize := len(queue)
		for k := 0; k < tmpSize; k++ {
			bad := queue[0]
			queue = queue[1:]
			for i := 0; i < 4; i++ {
				sr := bad[0] + dist[i][0]
				sc := bad[1] + dist[i][1]
				if sr < 0 || sc < 0 || sr >= len(grid) || sc >= len(grid[0]) || grid[sr][sc] != 1 {
					continue
				}

				queue = append(queue, [2]int{sr, sc})
				scan--
				grid[sr][sc] = 2
			}
		}

	}

	if scan > 0 {
		return -1
	}
	return ret - 1
}
