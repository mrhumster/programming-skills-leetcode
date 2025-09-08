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
	for i := 0; i < len(nums); i++ {
		j := slices.Index(nums, k-nums[i])
		if j == -1 {
			nums = slices.Delete(nums, i, i+1)
			continue
		}
		nums = slices.Delete(nums, j, j+1)
		nums = slices.Delete(nums, i-1, i)
		count++
		i = 0
	}
	return count
}
