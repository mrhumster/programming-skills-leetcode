/*
Package t724

Given an array of integers nums, calculate the pivot index of this array.

The pivot index is the index where the sum of all the numbers strictly
to the left of the index is equal to the sum of all the numbers strictly
to the index's right.

If the index is on the left edge of the array, then the left sum is 0
because there are no elements to the left. This also applies to the right
edge of the array.

Return the leftmost pivot index. If no such index exists, return -1.
*/
package t724

func pivotIndex(nums []int) int {
	lsum, rsum := 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		rsum += nums[i]
	}
	for i := range nums {
		rsum -= nums[i]
		if lsum == rsum {
			return i
		}
		lsum += nums[i]
	}
	return -1
}
