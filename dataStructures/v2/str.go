package v2

import (
	"fmt"
	"strings"
)

func urlify(word string) string {
	splitWord := strings.Split(word, "")

	for start, end := 0, len(splitWord)-1; start < end; start, end = start+1, end-1 {
		if splitWord[start] == " " {
			splitWord[start] = "%20"
		}
		if splitWord[end] == " " {
			splitWord[end] = "%20"
		}
	}

	return strings.Join(splitWord, "")
}

// Palindrome Permutation.
func isPalindromePermutation(phrase string) bool {
	// assumption of lowercase ASCII characters.
	charList := charCount(strings.ToLower(phrase))
	return isOneMaxOdd(charList)
}

func charCount(phrase string) [25]int32 {
	// the characters are between a - z
	var chars [25]int32

	for _, char := range phrase {
		if char >= 'a' && char <= 'z' {
			index := char - 'a'
			chars[index]++
		}
	}
	return chars
}

func isOneMaxOdd(chars [25]int32) bool {
	var found bool

	for _, count := range chars {
		if count%2 == 1 {
			if found {
				return false
			}
			found = true
		}
	}
	return true
}

// One Away
func oneAway(smallestStr, largestStr string) bool {
	if len(smallestStr) > len(largestStr) {
		temp := largestStr
		largestStr = smallestStr
		smallestStr = temp
	}

	// edit one away should have length difference of 1 at most.
	editCount := len(largestStr) - len(smallestStr)
	if editCount > 1 {
		return false
	}

	charMapOfLargest := charMap(largestStr)
	for _, char := range smallestStr {
		if val, ok := charMapOfLargest[char]; !ok || val == 0 {
			editCount += 1
		} else {
			charMapOfLargest[char] -= 1
		}

		if editCount > 1 {
			return false
		}
	}

	return true
}

func charMap(str string) map[rune]int {
	record := make(map[rune]int)
	for _, char := range str {
		record[char]++
	}

	return record
}

// Rabin-Karp substring search algorithm
func hashCode(char byte) int {
	if char >= 'a' && char <= 'z' {
		return int(char-'a') + 1
	}
	return 0
}

func hash(str string) int {
	power, sum := float64(0), 0

	for i := len(str) - 1; i >= 0; i -= 1 {
		sum += hashCode(str[i])
		power += 1
	}
	return sum
}

func searchSubString(sentence, subStr string) int {
	subStrHash := hash(subStr)
	preHash, count := hash(sentence[0:len(subStr)]), 0

	if subStrHash == preHash {
		count += 1
	}

	for prev, next := 1, len(subStr); next < len(sentence); prev, next = prev+1, next+1 {
		currHash := (preHash - hashCode(sentence[prev-1])) + hashCode(sentence[next])

		if currHash == subStrHash {
			count += 1
		}
		fmt.Println("match: ", subStrHash, "-", currHash, "-", prev, next)
		preHash = currHash

	}

	fmt.Println(len(sentence))

	return count
}

// compress string.
func compressString(str string) string {
	var sb strings.Builder

	for i := 0; i < len(str); {
		freq := 1

		for j := i + 1; j < len(str); j += 1 {
			if str[i] == str[j] {
				freq += 1
				continue
			}
			break
		}

		compstr := fmt.Sprintf("%c%d", str[i], freq)
		sb.WriteString(compstr)
		i += freq
	}

	if len(sb.String()) > len(str) {
		return str
	}
	return sb.String()
}

// string rotation - 1
func stringRotationV1(phrase, str string) bool {
	if len(phrase) == 0 || len(str) > len(phrase) {
		return false
	}

	strHash := hashB(str)
	prevHash := hashB(phrase[0:len(str)])

	if strHash == prevHash {
		return true
	}

	for start, end := 0, len(str); end < len(phrase); start, end = start+1, end+1 {
		currHash := (prevHash - hashB(string(phrase[start]))) + hashB(string(phrase[end]))
		if strHash == currHash {
			return true
		}
		prevHash = currHash
	}
	return false
}

func hashB(str string) int32 {
	sum := int32(0)
	for _, char := range str {
		sum += char
	}
	return sum
}

// Given a string, find the length of the longest substring in it with no more than K distinct characters.
//
// You can assume that K is less than or equal to the length of the given string.
func longestWithUniqueChar(str string, k int) string {
	longestSubStr := ""

	for start, end := 0, k; end < len(str); end++ {
		if !isUniqueCharCount(str[start:end+1], k) {
			if (end - start) >= len(longestSubStr) {
				longestSubStr = str[start:end]
			}
			start += 1
		}

	}

	return longestSubStr
}

func isUniqueCharCount(substr string, k int) bool {
	charMap := make(map[rune]bool) // Assume ASCII characters
	count := 0

	for _, char := range substr {
		if _, ok := charMap[char]; !ok {
			charMap[char] = true
			count += 1
		}
	}

	return count <= k
}

//////////////////////////////////////////////////////////////
//func main() {
//	fmt.Printf("%s: ", maxSubArrUniqueChar([]byte{'A', 'B', 'C', 'B', 'B', 'C', 'E'}, 3))
//}
// You are visiting a farm to collect fruits. The farm has a single row of fruit trees. You will be given two baskets, and your goal is to pick as many fruits as possible to be placed in the given baskets.
//
//You will be given an array of characters where each character represents a fruit tree. The farm has following restrictions:
//
//Each basket can have only one type of fruit. There is no limit to how many fruit a basket can hold.
//You can start with any tree, but you can’t skip a tree once you have started.
//You will pick exactly one fruit from every tree until you cannot, i.e., you will stop when you have to pick from a third fruit type.
//Write a function to return the maximum number of fruits in both baskets.
/////////////////

func maxSubArrUniqueChar(arr []byte, k int) []byte {
	charMap := make(map[byte]bool)
	var subStr []byte
	start, end := 0, 0

	for end < len(arr) {
		charMap[arr[end]] = true
		fmt.Printf("start: %+v\n", charMap)
		if len(charMap) <= k {
			end += 1

			if len(subStr) < len(arr[start:end]) {
				subStr = arr[start:end]
			}

		} else {
			delete(charMap, arr[start])
			start += 1
		}
		fmt.Printf("end: %+v\n", charMap)
	}

	return subStr
}

///////////////////////////////////////////////////////
//func main() {
//	fmt.Printf("%s: ", longestWithDistinctChars([]byte{'a', 'a', 'b', 'c', 'c', 'b', 'b'}))
//}
// Given a string, find the length of the longest substring, which has all distinct characters.

func longestWithDistinctChars(arr []byte) []byte {
	charMap := make(map[byte]bool)
	var subStr []byte
	start, end := 0, 0

	for end < len(arr) {
		fmt.Printf("start: %+v\n", charMap)
		if _, ok := charMap[arr[end]]; !ok {
			charMap[arr[end]] = true
			end += 1

			if len(subStr) <= len(arr[start:end]) {
				subStr = arr[start:end]
			}

		} else {
			delete(charMap, arr[start])
			start += 1
		}
		fmt.Printf("end: %+v\n", charMap)
	}

	return subStr
}
