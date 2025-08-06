/*
Package t58

Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal substring consisting of non-space characters only.
*/
package t58

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	arr := strings.Split(s, " ")
	word := strings.Trim(arr[len(arr)-1], " ")
	return len(word)
}
