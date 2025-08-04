package t242

func isAnagram(s string, t string) bool {
	/*
			 * An anagram is a word or phrase formed by rearranging the
		   * letters of a different word or phrase, using all the original l
			 * etters exactly once.
	*/
	if len(s) != len(t) {
		return false
	}
	charCount := make(map[byte]int)
	for i := range s {
		charCount[s[i]]++
	}
	for i := range t {
		charCount[t[i]]--
	}
	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}
	return true
}
