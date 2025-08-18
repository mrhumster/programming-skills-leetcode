/*
Package t976
Given an integer array nums, return the largest perimeter of a
triangle with a non-zero area, formed from three of these lengths.
If it is impossible to form any triangle of a non-zero area, return 0.
*/
package t976

import "sort"

func isATrianglePossible(a int, b int, c int) bool {
	return a < b+c
}

func calculatePerimeter(a int, b int, c int) int {
	return a + b + c
}

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		a := nums[i]
		b := nums[i-1]
		c := nums[i-2]
		if isATrianglePossible(a, b, c) {
			return calculatePerimeter(a, b, c)
		}
	}
	return 0
}
