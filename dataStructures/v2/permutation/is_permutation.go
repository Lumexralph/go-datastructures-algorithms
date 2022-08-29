package permutation

func isPermutation(firstWord, secondWord string) bool {
	if firstWord == secondWord {
		return false
	}
	if len(firstWord) != len(secondWord) {
		return false
	}

	sumA := sumCodePoints(firstWord)
	sumB := sumCodePoints(secondWord)

	return sumA == sumB
}

func sumCodePoints(word string) int32 {
	var sum int32

	for _, i := range word {
		sum += i
	}
	return sum
}
