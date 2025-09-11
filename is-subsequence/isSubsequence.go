/*
Package t392
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
Дано 2 строки s и t, вернуть true если s это подстрока t или false в другом случае.

A subsequence of a string is a new string that is formed from the original string by
deleting some (can be none) of the characters without disturbing the relative positions
of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
*/
package t392

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	is := 0
	for it := 0; it < len(t); it++ {
		if is >= len(s) {
			return true
		}
		if s[is] == t[it] {
			is++
		}

	}
	return is == len(s)
}
