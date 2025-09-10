/*
Package t1679

You are given an integer array nums and an integer k.

In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

Return the maximum number of operations you can perform on the array.
*/
package t1679

import "slices"

func maxOperations(nums []int, k int) int {
	count := 0
	slices.Sort(nums)
	for start, end := 0, len(nums)-1; start<end; {
		a := nums[start]
		b := nums[end]
		if a + b == k {
			start++
			end--
			count++
		} else if a + b < k {
			start++
		} else {
			end--
		}
	}
	return count
}
