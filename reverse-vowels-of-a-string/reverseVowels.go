/*
Package t345
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.
*/
package t345

import "slices"

func reverseVowels(s string) string {
	runes := []rune(s)
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	left, right := 0, len(runes)-1
	for left < right {
		for left < right && !slices.Contains(vowels, runes[left]) {
			left++
		}
		for left < right && !slices.Contains(vowels, runes[right]) {
			right--
		}
		if left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}
	return string(runes)
}
