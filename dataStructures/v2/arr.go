package v2

import (
	"math"
)

// Sliding Window Pattern Maximum Average

func maxAverageSubArr(arr []int, length int) float64 {
	// calculate sum of first sub-array 0 - length
	if length <= 0 || length > len(arr) {
		return 0
	}

	sum := sumSubArr(arr[:length])
	avg := math.Max(-math.MaxFloat64, sum/float64(length))

	for start, end := 0, length; end < len(arr); start, end = start+1, end+1 {
		prev := arr[start]
		next := arr[end]
		newSum := sum + float64(next-prev)
		avg = math.Max(avg, newSum/float64(length))
	}

	return avg
}

func sumSubArr(subArr []int) float64 {
	sum := 0
	for _, elem := range subArr {
		sum += elem
	}
	return float64(sum)
}

// Longest Substring Without Repeating Characters.
func longestSubStr(word string) int {
	var count, tempCount float64

	record := make(map[rune]bool)
	for _, char := range word {
		if exist := record[char]; exist {
			record[char] = false
			count = math.Max(count, tempCount)
			tempCount = 0
		}
		tempCount += 1
		record[char] = true
	}

	return int(count)
}

// rotate image 90 degrees
func rotateImage(img [][]int) [][]int {
	// validation for empty img
	// img with just one element arrays

	// to avoid pointing to same array during mutation
	tempTop := make([]int, len(img[0]))
	copy(tempTop, img[0])
	imgLength := len(img)

	// left ==> top
	for i := 0; i < imgLength; i += 1 {
		img[0][i] = img[(imgLength-1)-i][0]
	}

	// bottom ==> left
	for i := 0; i < imgLength; i += 1 {
		img[i][0] = img[imgLength-1][i]
	}

	// right ==> bottom
	for i := 0; i < imgLength; i += 1 {
		img[imgLength-1][i] = img[(imgLength-1)-i][(imgLength - 1)]
	}

	// top ==> right
	for i := 0; i < imgLength; i++ {
		img[i][imgLength-1] = tempTop[i]
	}

	return img
}

// Zero Matrix
func zeroMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}

	colRecord := make(map[int]struct{})

	for row := 0; row < len(matrix); row += 1 {
		for col := 0; col < len(matrix[row]); col += 1 {
			if _, ok := colRecord[col]; ok {
				continue
			}
			if matrix[row][col] == 0 {
				colRecord[col] = struct{}{}

				// blank out the row to 0
				for i := 0; i < len(matrix[row]); i += 1 {
					matrix[row][i] = 0
				}
				// blank out the column
				for j := 1; j < len(matrix); j += 1 {
					matrix[j][col] = 0
				}
				break
			}
		}
	}

	return matrix
}

func binarySearch(arr []int, item int) bool {
	if len(arr) == 0 {
		return false
	}

	if len(arr) == 1 {
		return arr[0] == item
	}

	start := 0
	end := len(arr) - 1
	mid := (start + end) / 2

	for start <= end {
		val := arr[mid]
		if item == val {
			return true
		} else if val > item {
			end = mid - 1
		} else {
			start = mid + 1
		}

		mid = (start + end) / 2
	}

	return false
}

func searchRotatedArr(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 && nums[0] == target {
		return 0
	}

	mid := len(nums) / 2
	leftIndex := mid
	rightIndex := mid

	// nums of an even number length, for evenly spread of index
	if len(nums)%2 == 0 {
		leftIndex = mid - 1
	}

	for leftIndex >= 0 && rightIndex < len(nums) {
		if nums[leftIndex] == target {
			return leftIndex
		}
		if nums[rightIndex] == target {
			return rightIndex
		}

		rightIndex += 1
		leftIndex -= 1
	}

	return -1
}
