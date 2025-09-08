/*
Package t334

Given an integer array nums, return true if there exists a
triple of indices (i, j, k) such that i < j < k and
nums[i] < nums[j] < nums[k]. If no such indices exists,
return false.

Если задан целочисленный массив nums, верните значение true,
если существует тройка индексов (i, j, k), таких что i < j < k и
nums[i] < nums[j] < nums[k]. Если таких индексов не существует,
верните значение false.
*/
package t334

import "math"

func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt
	for _, n := range nums {
		if n <= first {
			first = n
		} else if n <= second {
			second = n
		} else {
			return true
		}
	}
	return false
}
