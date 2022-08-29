package isunique

func isUnique(word string) bool {
	for i, j := 0, len(word)-1; i < len(word) && j >= 0; i, j = i+1, j-1 {
		if word[i] == word[j] && i != j {
			return false
		}
	}
	return len(word) > 0
}

func isUniqueV2(word string) bool {
	store := make(map[uint8]bool)

	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if found := store[word[i]]; found {
			return false
		}
		store[word[i]] = true

		if found := store[word[j]]; found {
			return false
		}
		store[word[j]] = true
	}
	return len(word) > 0
}
