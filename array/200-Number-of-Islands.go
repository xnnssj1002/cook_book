package array

import (
	"container/list"
)

/* 题目：200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
*/

/* 题目大意
给定一个由 ‘1’（陆地）和 ‘0’（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
*/

/* 示例
示例 1：所有的陆地连在一起，所以只有一个岛屿
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1

示例 2：有三块独立的陆地，所以有三个岛屿
输入：var grid = [][]byte{
	[]byte{'1', '1', '0', '0', '0'},
	[]byte{'1', '1', '0', '0', '0'},
	[]byte{'0', '0', '1', '0', '0'},
	[]byte{'0', '0', '0', '1', '1'},
}
输出：3
*/

// ********** 深度优先 ********** //

// NumIslandsByDfs 深度优先实现
// 为了求出岛屿的数量，我们可以遍历整个二维网格。如果一个位置为 '1' ，则以其为起始节点开始进行深度优先搜索。
// 在深度优先搜索的过程中，每个搜索到的 '1' 都会被重新标记为 '0'。
// 最终岛屿的数量就是我们进行深度优先搜索的次数。
func NumIslandsByDfs(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}
func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0' //避免重复遍历
	// 对给出的节点[i, j]的 上下左右 四个节点，分别进行搜索，即做到深度优先
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}

// NumIslandsByDfsSelf 根据题解，自己实现深度搜索
func NumIslandsByDfsSelf(grid [][]byte) int {
	count := 0

	var dfs func([][]byte, int, int)
	dfs = func(grids [][]byte, i int, j int) {
		// 判断i, j 不能小于0， 不能大于对应的长度，且[i, j]对应的元素必须是 字节'1'
		if i < 0 || j < 0 || i >= len(grids) || j >= len(grids[0]) || grids[i][j] != '1' {
			return
		}
		// 将[i, j]对应的元素改为'0'，防止重复处理
		grids[i][j] = '0'
		// 对该元素[i, j]的上下左右的元素进行搜索
		dfs(grids, i, j-1) // 上
		dfs(grids, i, j+1) // 下
		dfs(grids, i-1, j) // 左
		dfs(grids, i+1, j) // 右
	}

	// 对grid[i][j]==1的元素，进行深度搜索
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}

	return count
}

// ********** 广度优先 ********** //

// NumIslandsBfs 广度优先实现
// 别人实现的代码
func NumIslandsBfs(grid [][]byte) int {
	count := 0
	optionX := []int{-1, 0, 0, 1}
	optionY := []int{0, 1, -1, 0}
	var queue [][]int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				queue = append(queue, []int{i, j})
				grid[i][j] = '0'
				for len(queue) != 0 {
					u := queue[0]
					queue = queue[1:]
					for m := 0; m < 4; m++ {
						newX := u[0] + optionX[m]
						newY := u[1] + optionY[m]
						if newX >= 0 && newY >= 0 && newX < len(grid) && newY < len(grid[0]) && grid[newX][newY] == '1' {
							queue = append(queue, []int{newX, newY})
							grid[newX][newY] = '0'
						}
					}
				}
			}
		}
	}
	return count
}

// NumIslandsByBfsSelf 自己实现广度搜索
// 为了求出岛屿的数量，我们可以遍历整个二维网格。如果一个位置为 '1'，则将其加入队列，开始进行广度优先搜索。
// 在广度优先搜索的过程中，每个搜索到的 '1' 都会被重新标记为 '0'。直到队列为空，搜索结束。
// 最终岛屿的数量就是我们进行广度优先搜索的次数。
func NumIslandsByBfsSelf(grid [][]byte) int {
	type item struct {
		row int
		col int
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				grid[i][j] = '0'
				queue := list.New()
				queue.PushBack(item{row: i, col: j})
				for queue.Len() > 0 {
					temp := queue.Front()
					queue.Remove(temp)
					tempItem := temp.Value.(item)

					// 左节点
					if tempItem.col-1 >= 0 && grid[tempItem.row][tempItem.col-1] == '1' {
						grid[tempItem.row][tempItem.col-1] = '0'
						queue.PushBack(item{row: tempItem.row, col: tempItem.col - 1})
					}
					// 右节点
					if tempItem.col+1 < len(grid[0]) && grid[tempItem.row][tempItem.col+1] == '1' {
						grid[tempItem.row][tempItem.col+1] = '0'
						queue.PushBack(item{row: tempItem.row, col: tempItem.col + 1})
					}
					// 上节点
					if tempItem.row-1 >= 0 && grid[tempItem.row-1][tempItem.col] == '1' {
						grid[tempItem.row-1][tempItem.col] = '0'
						queue.PushBack(item{row: tempItem.row - 1, col: tempItem.col})
					}
					// 下节点
					if tempItem.row+1 < len(grid) && grid[tempItem.row+1][tempItem.col] == '1' {
						grid[tempItem.row+1][tempItem.col] = '0'
						queue.PushBack(item{row: tempItem.row + 1, col: tempItem.col})
					}

				}
			}
		}
	}

	return count
}
