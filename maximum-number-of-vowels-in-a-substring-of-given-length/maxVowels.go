/*
Package t1456

Given a string s and an integer k, return the maximum number of
vowel letters in any substring of s with length k.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.
*/
package t1456


func isVowel(ch byte) bool {
	return  ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func maxVowels(s string, k int) int {
	maxVowelsSum := 0
	for i := range k {
		if isVowel(s[i]) {
			maxVowelsSum++
		}
	}
	sum := maxVowelsSum
	for start, end := 1, k; end<len(s); start, end = start + 1, end + 1 {
		if isVowel(s[start-1]) {
			sum--
		}
		if isVowel(s[end]) {
			sum++
		}
		if sum > maxVowelsSum {
			maxVowelsSum = sum 
		}
	}
	return maxVowelsSum
}

