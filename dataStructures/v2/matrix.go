package v2

import (
	"fmt"
	"strings"
)

/////////////////////////////////////////////////////////////////////////////
//func main() {
//	land := [][]int{
//		{0, 1, 1, 1, 0},
//		{0, 0, 0, 1, 1},
//		{0, 1, 1, 1, 0},
//		{0, 1, 1, 0, 0},
//		{0, 0, 0, 0, 0},
//	}
//	fmt.Printf("islands: %d", countIslandsDFS(land))
//}

func countIslandsDFS(matrix [][]int) int {
	count, row, col := 0, len(matrix), len(matrix[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 1 {
				count += 1
				visitIsland(matrix, i, j)
			}
		}
	}

	return count
}

func visitIsland(matrix [][]int, i, j int) {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return
	}

	if matrix[i][j] == 0 {
		return
	}
	matrix[i][j] = 0 // mark the cell

	// Vertical downwards / upwards
	visitIsland(matrix, i+1, j)
	visitIsland(matrix, i-i, j)

	// <== Horizontal ==>
	visitIsland(matrix, i, j+1)
	visitIsland(matrix, i, j-1)
}

/////////////////////////////////////////////////////////////////////////////////////
//func main() {
//	land := [][]int{
//		{1, 1, 1, 0, 0},
//		{0, 1, 0, 0, 1},
//		{0, 0, 1, 1, 0},
//		{0, 1, 1, 0, 0},
//		{0, 0, 1, 0, 0},
//	}
//	fmt.Printf("biggest-islands: %d", biggestIslandsDFS(land))
//}

func biggestIslandsDFS(matrix [][]int) int {
	count, row, col := 0, len(matrix), len(matrix[0])
	biggestArea := 0

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 1 {
				count += 1
				area := countIsland(matrix, i, j, 0)
				if area > biggestArea {
					biggestArea = area
				}
			}
		}
	}

	return biggestArea
}

func countIsland(matrix [][]int, i, j, area int) int {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return area
	}

	if matrix[i][j] == 0 {
		return area
	}
	matrix[i][j] = 0 // mark the cell
	area += 1

	// Vertical downwards / upwards
	c1 := countIsland(matrix, i+1, j, area)
	c2 := countIsland(matrix, i-i, j, c1)

	// <== Horizontal ==>
	c3 := countIsland(matrix, i, j+1, c2)
	c4 := countIsland(matrix, i, j-1, c3)

	return c4
}

///////////////////////////////////////////////////////////////////////////////
//func main() {
//	image := [][]int{
//		{0, 1, 1, 1, 0},
//		{0, 0, 0, 1, 1},
//		{0, 1, 1, 1, 0},
//		{0, 1, 1, 0, 0},
//		{0, 0, 0, 0, 0},
//	}
//
//	floodFill(image, 1, 3, image[1][3], 2)
//	fmt.Printf("filled: %+v", image)
//}

func floodFill(matrix [][]int, row, col, val, colour int) {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
		return
	}

	if matrix[row][col] != val {
		return
	}
	matrix[row][col] = colour // mark the cell

	// Vertical downwards / upwards
	floodFill(matrix, row+1, col, val, colour)
	floodFill(matrix, row-1, col, val, colour)

	// <== Horizontal ==>
	floodFill(matrix, row, col+1, val, colour)
	floodFill(matrix, row, col-1, val, colour)
}

////////////////////////////////////////////////////////////////////////////
//func main() {
//	layer := [][]int{
//		{0, 0, 0, 0},
//		{0, 1, 0, 0},
//		{0, 1, 0, 0},
//		{0, 0, 1, 0},
//		{0, 0, 0, 0},
//	}
//
//	fmt.Printf("closed-islands: %+v", closedIsland(layer))
//}

func closedIsland(matrix [][]int) int {
	closed := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				if closedIslandDFS(matrix, i, j) {
					closed += 1
				}
			}
		}
	}

	return closed
}

func closedIslandDFS(matrix [][]int, row, col int) bool {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[row]) {
		return false
	}

	if matrix[row][col] == 0 || matrix[row][col] == 2 {
		return true
	}
	matrix[row][col] = 2

	// Vertical upwards / downwards
	c1 := closedIslandDFS(matrix, row-1, col) // upwards
	c2 := closedIslandDFS(matrix, row+1, col) // downwards

	// Horizontal <== left right ==>
	c3 := closedIslandDFS(matrix, row, col-1) // left
	c4 := closedIslandDFS(matrix, row, col+1) // right

	return c1 && c2 && c3 && c4
}

////////////////////////////////////////////////////////////////////////
//func main() {
//	layer := [][]int{
//		{0, 0, 0, 0},
//		{0, 1, 0, 0},
//		{0, 1, 0, 0},
//		{0, 1, 1, 0},
//		{0, 1, 0, 0},
//	}
//
//	fmt.Printf("perimeter: %d", islandPerimeter(layer))
//}

func islandPerimeter(layer [][]int) int {
	lands := 0

	for row := 0; row < len(layer); row++ {
		for col := 0; col < len(layer[row]); col++ {
			if layer[row][col] == 1 {
				lands = countLandDFS(layer, row, col)
			}
		}
	}
	// ((lands - 2) * 2) + 6.. a trick
	return lands
}

func countLandDFS(layer [][]int, row, col int) int {
	if row < 0 || row >= len(layer) || col < 0 || col >= len(layer[row]) {
		return 1
	}

	if layer[row][col] == 0 {
		return 1
	}

	if layer[row][col] == 2 {
		return 0
	}
	layer[row][col] = 2

	count := 0
	// Vertical count (upwards, downwards)
	count += countLandDFS(layer, row-1, col)
	count += countLandDFS(layer, row+1, col)

	// Horizontal (left, right)
	count += countLandDFS(layer, row, col-1)
	count += countLandDFS(layer, row, col+1)

	return count
}

/////////////////////////////////////////////////////////////////////////////////
//func main() {
//	layer := [][]int{
//		{1, 1, 0, 1, 1},
//		{1, 1, 0, 1, 1},
//		{0, 0, 0, 0, 0},
//		{0, 1, 1, 0, 1},
//		{0, 1, 1, 0, 1},
//	}
//
//	fmt.Printf("unique-islands: %d", countUniqueIslands(layer))
//}

func countUniqueIslands(layer [][]int) int {
	uniqueIslands := make(map[string]struct{})

	for row := 0; row < len(layer); row++ {
		for col := 0; col < len(layer[row]); col++ {
			pattern := islandPattern(layer, row, col)
			if pattern != "" {
				uniqueIslands[pattern] = struct{}{}
			}
		}
	}
	fmt.Printf("%+v", uniqueIslands)
	return len(uniqueIslands)
}

func islandPattern(layer [][]int, row, col int) string {
	if row < 0 || row >= len(layer) || col < 0 || col >= len(layer[row]) {
		return ""
	}

	if layer[row][col] == 0 {
		return ""
	}
	layer[row][col] = 0

	var pattern strings.Builder

	// Vertical (upwards, downwards)
	pattern.WriteString("U" + islandPattern(layer, row-1, col))
	pattern.WriteString("D" + islandPattern(layer, row+1, col))

	// Horizontal (left, right)
	pattern.WriteString("L" + islandPattern(layer, row, col-1))
	pattern.WriteString("R" + islandPattern(layer, row, col+1))

	return pattern.String()

}
