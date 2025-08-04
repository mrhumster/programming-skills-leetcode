package t389

func findTheDifference(s string, t string) byte {
	charCount := make(map[byte]int)

	for i := range t {
		charCount[t[i]]++
	}

	for i := range s {
		charCount[s[i]]--
	}

	for char, count := range charCount {
		if count > 0 {
			return char
		}
	}

	return 0
}
