package v2

import (
	"math"
	"sort"
)

// Input: [2, 3, 3, 3, 6, 9, 9]
//Output: 4
//Explanation: The first four elements after removing the duplicates will be [2, 3, 6, 9].

//func main() {
//	arr := []int{2, 2, 2, 11}
//
//	fmt.Printf("unique: %+v", removeDuplicatesA(arr))
//}

func removeDuplicatesA(arr []int) []int {
	start := 0
	for end := 1; end < len(arr); end++ {
		if arr[start] != arr[end] {
			start += 1
			val := arr[end]
			arr[start] = val
		}
	}

	return arr[:start+1]
}

///////////////////////////////////////////////////////////////////////
//func main() {
//	arr := []int{3, 2, 3, 6, 3, 10, 9, 3}
//	fmt.Printf("%+v", removeKeyInstance(arr, 3))

//Input: [3, 2, 3, 6, 3, 10, 9, 3], Key=3
//Output: 4
//Explanation: The first four elements after removing every 'Key' will be [2, 6, 10, 9].
//}

func removeKeyInstance(arr []int, key int) []int {
	start := 0
	for end := 0; end < len(arr); end++ {
		if arr[end] != key {
			arr[start] = arr[end]
			start += 1
		}
	}

	return arr[:start]
}

//////////////////////////////////////////////////////////////
//func main() {
//	arr := []int{-2, -1, 0, 2, 3}
//	fmt.Printf("%+v", sortedSquares(arr))
//}

func sortedSquares(arr []int) []int {
	length := len(arr)
	squaredArr := make([]int, length, length)
	left, right := 0, length-1
	highestIndex := length - 1

	for left <= right {
		leftSquare := arr[left] * arr[left]
		rightSquare := arr[right] * arr[right]

		if leftSquare > rightSquare {
			squaredArr[highestIndex] = leftSquare
			left += 1
		} else {
			squaredArr[highestIndex] = rightSquare
			right -= 1
		}
		highestIndex -= 1
	}

	return squaredArr
}

/////////////////////////////////////////////////////////
/*
Input: [-3, 0, 1, 2, -1, 1, -2]
Output: [-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]
Explanation: There are four unique triplets whose sum is equal to zero.
*/
//func main() {
//	arr := []int{-3, 0, 1, 2, -1, 1, -2}
//	fmt.Printf("triplets: %+v\n", uniqueTriplets(arr))
//
//	arr = []int{-5, 2, -1, -2, 3}
//	fmt.Printf("triplets: %+v\n", uniqueTriplets(arr))
//}

func uniqueTriplets(arr []int) [][]int {
	// sort to make duplicates close to be easily skipped
	sort.IntSlice(arr).Sort()

	var triplets [][]int
	for i := 0; i < len(arr); i++ {
		if i > 0 && arr[i] == arr[i-1] {
			// skip duplicates
			continue
		}
		triplets = append(triplets, findPairs(arr, -arr[i], i+1)...)
	}

	return triplets
}

func findPairs(arr []int, target, startIndex int) [][]int {
	endIndex := len(arr) - 1

	var pairs [][]int

	for startIndex < endIndex {
		currSum := arr[startIndex] + arr[endIndex]
		if currSum == target {
			pairs = append(pairs, []int{-target, arr[startIndex], arr[endIndex]})
			startIndex += 1
			endIndex -= 1

			// skip duplicates
			for startIndex < endIndex && arr[startIndex] == arr[startIndex-1] {
				startIndex += 1
			}
			for startIndex < endIndex && arr[endIndex] == arr[endIndex+1] {
				endIndex -= 1
			}
		} else if currSum > target { // left will be larger
			endIndex -= 1
		} else {
			startIndex += 1
		}
	}
	return pairs
}

///////////////////////////////////////////////////////////////
/*func main() {
	arr := []int{2, 5, 3, 10}
	fmt.Printf("%+v\n", productSubArr(arr, 30))

	arr = []int{8, 2, 6, 5}
	fmt.Printf("%+v\n", productSubArr(arr, 50))
}

Input: [8, 2, 6, 5], target=50
Output: [8], [2], [8, 2], [6], [2, 6], [5], [6, 5]
Explanation: There are seven contiguous subarrays whose product is less than the target.
*/

func productSubArr(arr []int, product int) [][]int {
	var productArr [][]int

	for i := 0; i < len(arr); i++ {
		if arr[i] < product {
			productArr = append(productArr, []int{arr[i]})

			right := i + 1
			if right == len(arr) {
				// end
				break
			}
			val := arr[i] * arr[right]
			for right < len(arr) {
				if val >= product {
					break
				}
				productArr = append(productArr, arr[i:right+1])
				right += 1
				if right != len(arr) {
					val = arr[right] * val
				}

			}
		}
	}
	return productArr
}

///////////////////////////////////////////////////////////////////////////////
/*
Given an array containing 0s, 1s and 2s, sort the array in-place.
Input: [2, 2, 0, 1, 2, 0]
Output: [0 0 1 2 2 2 ]

func main() {
	arr := []int{1, 0, 2, 1, 0}
	sortArray(arr)
	fmt.Printf("%+v\n", arr)

	arr = []int{2, 2, 0, 1, 2, 0}
	sortArray(arr)
	fmt.Printf("%+v\n", arr)
}
*/

func sortArray(arr []int) {
	low, high := 0, len(arr)-1

	for i := 0; i <= high; {
		if arr[i] == 0 {
			temp := arr[low]
			arr[low] = arr[i]
			arr[i] = temp

			low += 1
			i += 1
		} else if arr[i] == 1 {
			i += 1
		} else {
			temp := arr[high]
			arr[high] = arr[i]
			arr[i] = temp
			high -= 1
		}
	}
}

// ////////////////////////////////////////////////////////////////////
type interval struct {
	start, end int
}

/*
func main() {
	intervals := []interval{
		interval{start: 6, end: 7},
		interval{start: 2, end: 4},
		interval{start: 5, end: 9},
	}
	fmt.Printf("merged-interval: %+v", mergeIntervals(intervals))
}

Intervals: [[6,7], [2,4], [5,9]]
Output: [[2,4], [5,9]]
Explanation: Since the intervals [6,7] and [5,9] overlap, we merged them into one [5,9].
*/

func mergeIntervals(intervals []interval) []interval {
	if len(intervals) < 2 {
		return intervals
	}

	// sort intervals by start
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	var merged []interval
	start := intervals[0].start
	end := intervals[0].end

	for _, interv := range intervals[1:] {
		if interv.start <= end { // there's an overlap
			end = int(math.Max(float64(end), float64(interv.end)))
		} else {
			merged = append(merged, interval{start, end})
			// reset
			start = interv.start
			end = interv.end
		}
	}
	merged = append(merged, interval{start, end})

	return merged
}

/////////////////////////////////////////////////////////////////////////
// Incomplete - Intervals
/*
func main() {
	intervals := []interval{
		interval{start: 1, end: 3},
		interval{start: 5, end: 7},
		interval{start: 8, end: 12},
	}
	fmt.Printf("merged-interval: %+v", mergeIntervalsWithInterval(intervals, interval{start: 4, end: 6}))
}
*/

func mergeIntervalsWithInterval(intervals []interval, interv interval) []interval {
	var merged []interval

	i := 0
	for i < len(intervals) {
		currInterv := intervals[i]
		if interv.start <= currInterv.end { // overlaps
			start := int(math.Min(float64(interv.start), float64(currInterv.start)))
			end := interv.end

			for i < len(intervals) && interv.end <= intervals[i].end {
				end = intervals[i].end
				i += 1
			}

			merged = append(merged, interval{start, end})
		} else {
			merged = append(merged, currInterv)
			i += 1
		}
	}

	return merged
}

////////////////////////////////////////////////////////////
// Cyclic Sort
/*
Input: [2, 3, 1, 8, 2, 3, 5, 1]
Output: 4, 6, 7
Explanation: The array should have all numbers from 1 to 8, due to duplicates 4, 6, and 7 are missing.

func main() {
	sequences := []int{2, 3, 1, 8, 2, 3, 5, 1}
	sortSequences(sequences)
	fmt.Printf("seq: %+v\n", findMissingNumbers(sequences))

	sequences = []int{2, 4, 1, 2}
	sortSequences(sequences)
	fmt.Printf("seq: %+v\n", findMissingNumbers(sequences))

	sequences = []int{2, 3, 2, 1}
	sortSequences(sequences)
	fmt.Printf("seq: %+v\n", findMissingNumbers(sequences))
}

*/

func sortSequences(sequences []int) {
	i := 0
	for i < len(sequences) {
		j := sequences[i] - 1
		if sequences[i] != sequences[j] {
			sequences[i], sequences[j] = sequences[j], sequences[i]
		} else {
			i++
		}
	}
}

func findMissingNumbers(sequences []int) []int {
	var missing []int
	for i := 0; i < len(sequences); i++ {
		if sequences[i] != i+1 {
			missing = append(missing, i+1)
		}
	}

	return missing
}
