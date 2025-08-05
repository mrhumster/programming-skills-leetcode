/*
Package t459

Repeated Substring Pattern
Given a string s, check if it can be constructed by taking a substring
of it and appending multiple copies of the substring together

import "strings"

	func repeatedSubstringPattern(s string) bool {
	    double := s + s
	    return strings.Contains(double[1:len(double)-1], s)
	}
*/
package t459

func findRepeated(s string) (string, bool) {
	l := len(s)
	mn := 1
	mx := l / 2
	for i := mn; i <= mx; i++ {
		if l%i == 0 {
			ss := s[0:i]
			repeatCount := l / i
			expected := ""
			for j := 0; j < repeatCount; j++ {
				expected += ss
			}
			if expected == s {
				return ss, true
			}
		}
	}
	return "", false
}

func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}
	_, f := findRepeated(s)
	return f
}
