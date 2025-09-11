/*
Package t1004
Max Consecutive Ones III

Given a binary array nums and an integer k,
return the maximum number of consecutive 1's
in the array if you can flip at most k 0's.
*/
package t1004

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	maxLength := 0
	zeroCount := 0

	for right < len(nums) {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}

		right++
	}

	return maxLength
}
