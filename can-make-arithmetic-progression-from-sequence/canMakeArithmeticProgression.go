/*
Package t1502

## 1502. Can Make Arithmetic Progression From Sequence
A sequence of numbers is called an arithmetic progression if the difference between any two
consecutive elements is the same.

Given an array of numbers arr, return true if the array can be rearranged to form an arithmetic
progression. Otherwise, return false.
*/
package t1502

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	step := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] != step {
			return false
		}
	}
	return true
}
