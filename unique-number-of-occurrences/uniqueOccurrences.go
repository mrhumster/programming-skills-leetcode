/*
Package t1207

Given an array of integers arr, return true if
the number of occurrences of each value in the
array is unique or false otherwise.
*/
package t1207

import "slices"

func uniqueOiccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		m[v] += 1
	}
	e := []int{}
	for _, v := range m {
		if slices.Contains(e, v) {
			return false
		}
		e = append(e, v)
	}
	return true
}
