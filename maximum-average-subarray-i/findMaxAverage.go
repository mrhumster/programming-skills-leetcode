/*
Package t643

Вам задан целочисленный массив nums, состоящий из n элементов, и целое число k.

Найдите непрерывный подмассив, длина которого равна k и который имеет максимальное среднее
значение, и верните это значение. Будет принят любой ответ с ошибкой вычисления менее 10-5.

You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average
value and return this value. Any answer with a calculation error less than 10-5 will be accepted.
*/
package t643

import "slices"

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 1 {
		return float64(nums[0])
	}
	if k == 1 {
		return float64(slices.Max(nums))
	}
	sum := 0
	for i:=0; i<k; i++ {
		sum += nums[i]
	}
	if k == len(nums) {
		return float64(sum)/float64(k)
	}
	maxSum := sum
	for start, end := 1, k; end < len(nums); start, end = start+1, end+1 {
		
		sum = sum - nums[start-1] + nums[end] 
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum)/float64(k)
}
