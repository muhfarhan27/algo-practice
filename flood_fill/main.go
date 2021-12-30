// 733. Flood Fill
// Easy

// An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.

// You are also given three integers sr, sc, and newColor. You should perform a flood fill on the image starting from the pixel image[sr][sc].

// To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with newColor.

// Return the modified image after performing the flood fill.

// Example 1:

// Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, newColor = 2
// Output: [[2,2,2],[2,2,0],[2,0,1]]
// Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color.
// Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.

// Example 2:

// Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, newColor = 2
// Output: [[2,2,2],[2,2,2]]

package main

import "fmt"

func main() {
	var (
		image    = [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
		sr       = 1
		sc       = 1
		newColor = 2
	)
	fmt.Println(floodFill(image, sr, sc, newColor))
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	rows := len(image)
	cols := len(image[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	firstColor := image[sr][sc]
	if image[sr][sc] == newColor {
		return image
	} else {
		paint(image, newColor, sr, sc, rows, cols, directions, firstColor)
	}

	return image
}

func paint(image [][]int, newColor, row, col, rowLen, colLen int, directions [][]int, firstColor int) {
	if row < 0 || row >= rowLen || col < 0 || col >= colLen || image[row][col] != firstColor {
		return
	}
	if image[row][col] != newColor {
		image[row][col] = newColor
	}

	for _, direction := range directions {
		newRow := row + direction[0]
		newCol := col + direction[1]
		paint(image, newColor, newRow, newCol, rowLen, colLen, directions, firstColor)
	}
}
