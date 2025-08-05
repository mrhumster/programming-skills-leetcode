/*
Package t896

896. Monotonic Array
An array is monotonic if it is either monotone increasing or monotone decreasing.
Anarray nums is monotone increasing if for all i <= j, nums[i] <= nums[j]. An array nums is
monotone decreasing if for all i <= j, nums[i] >= nums[j].

Given an integer array nums, return true if the given array is monotonic, or false otherwise.
*/
package t896

func isMonotonic(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	direction := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		step := nums[i] - nums[i-1]
		if direction > 0 && step < 0 {
			return false
		}
		if direction < 0 && step > 0 {
			return false
		}
		if direction == 0 {
			direction = step
		}
	}
	return true
}
