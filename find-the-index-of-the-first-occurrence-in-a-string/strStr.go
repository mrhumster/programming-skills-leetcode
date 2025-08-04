package t28

/*
 * 28. Find the Index of the First Occurrence in a String
 * Given two strings needle and haystack, return the index of the first occurrence
 * of needle in haystack, or -1 if needle is not part of haystack.
 */

func strStr(heystack string, needle string) int {
	occurenceOfNeedle := -1
	for i, r := range heystack {
		if byte(r) != needle[0] {
			continue
		}
		j := len(needle) + i
		if j > len(heystack) {
			return occurenceOfNeedle
		}
		if heystack[i:j] == needle {
			return i
		}
	}
	return occurenceOfNeedle
}
