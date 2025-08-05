/*
Package p283

Given an integer array nums, move all 0's to the end of it
while maintaining the relative order of the non-zero elements.

**Note** that you must do this in-place without making a copy of the array.
*/
package p283

func moveZeroes(nums []int) []int {
	nonZeroIndex := 0 // Индекс для ненулевых элементов

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex] = nums[i]
			nonZeroIndex++
		}
	}

	for nonZeroIndex < len(nums) {
		nums[nonZeroIndex] = 0
		nonZeroIndex++
	}
	return nums
}
