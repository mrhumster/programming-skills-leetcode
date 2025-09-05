/*
Package t238

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.
*/
package t238

func productExceptSelf(nums []int) []int {
	out := make([]int, len(nums))
	out2 := make([]int, len(nums))
	nums1 := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			out[i] = 1
			continue
		}
		out[i] = out[i-1] * nums[i-1]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			out2[i] = nums[i]
			continue
		}
		out2[i] = out2[i+1] * nums[i]
	}
	nums1[0] = out2[1]
	nums1[len(nums1)-1] = out[len(nums1)-1]
	for i := 1; i < len(nums)-1; i++ {
		nums1[i] = out[i] * out2[i+1]
	}
	return nums1
}
