package twentyDays

/* 733. 图像渲染
https://leetcode-cn.com/problems/flood-fill/
*/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	colorDFS(image, sr, sc, oldColor, newColor)
	return image
}

func colorDFS(image [][]int, sr int, sc int, oldColor, newColor int) {
	if sc < 0 || sr < 0 || sr > len(image)-1 || sc > len(
		image[0])-1 || image[sr][sc] != oldColor || oldColor == newColor {
		return
	}

	image[sr][sc] = newColor
	colorDFS(image, sr+1, sc, oldColor, newColor)
	colorDFS(image, sr-1, sc, oldColor, newColor)
	colorDFS(image, sr, sc+1, oldColor, newColor)
	colorDFS(image, sr, sc-1, oldColor, newColor)
}

/* 695. 岛屿的最大面积
https://leetcode-cn.com/problems/max-area-of-island/
 */
func maxAreaOfIsland(grid [][]int) int {
	visited := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		visited[i] = make([]int, len(grid[0]))
	}

	retArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if visited[i][j] == 1 || grid[i][j] != 1 {
				continue
			}

			tmp := maxAreaOfIslandDFS(grid, visited, i, j, 0)
			if tmp > retArea {
				retArea = tmp
			}
		}
	}
	return retArea
}

func maxAreaOfIslandDFS(grid, visited [][]int, sr, sc, area int) int {
	if sc < 0 || sr < 0 || sr > len(grid)-1 || sc > len(grid[0])-1 || grid[sr][sc] != 1 || visited[sr][sc] == 1 {
		return area
	}

	area++
	visited[sr][sc] = 1
	area = maxAreaOfIslandDFS(grid, visited, sr+1, sc, area)
	area = maxAreaOfIslandDFS(grid, visited, sr-1, sc, area)

	area = maxAreaOfIslandDFS(grid, visited, sr, sc+1, area)
	area = maxAreaOfIslandDFS(grid, visited, sr, sc-1, area)
	return area
}
