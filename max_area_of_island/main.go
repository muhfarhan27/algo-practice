// 695. Max Area of Island
// Medium

// You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

// The area of an island is the number of cells with a value 1 in the island.

// Return the maximum area of an island in grid. If there is no island, return 0.

// Example 1:

// Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
// Output: 6
// Explanation: The answer is not 11, because the island must be connected 4-directionally.

// Example 2:

// Input: grid = [[0,0,0,0,0,0,0,0]]
// Output: 0

package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
	fmt.Println(maxAreaOfIsland2(grid))
}

func maxAreaOfIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	area := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			area = max(area, findIslands(grid, i, j, rows, cols, directions))
		}
	}
	return area
}

func findIslands(grid [][]int, row, col, rowLen, colLen int, directions [][]int) int {
	islands := 0
	if row < 0 || row >= rowLen || col < 0 || col >= colLen || grid[row][col] == 0 {
		return 0
	}
	grid[row][col] = 0
	islands++
	for _, direction := range directions {
		newRow := row + direction[0]
		newCol := col + direction[1]
		islands += findIslands(grid, newRow, newCol, rowLen, colLen, directions)
	}
	return islands
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//a bit tidier
func maxAreaOfIsland2(grid [][]int) int {
	area := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				area = max(area, findIslands2(grid, i, j))
			}
		}
	}
	return area
}

func findIslands2(grid [][]int, row, col int) int {
	islands := 0
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == 0 {
		return 0
	}
	grid[row][col] = 0
	islands++
	islands += findIslands2(grid, row-1, col)
	islands += findIslands2(grid, row+1, col)
	islands += findIslands2(grid, row, col+1)
	islands += findIslands2(grid, row, col-1)
	return islands
}
