/*
Package t1491

You are given an array of unique integers salary where salary[i]
is the salary of the ith employee.

Return the average salary of employees excluding the minimum and
maximum salary. Answers within 10-5 of the actual answer will be accepted.
*/
package t1491

import "slices"

func average(salary []int) float64 {
	slices.Sort(salary)
	sum := 0.0
	for i := 1; i < len(salary)-1; i++ {
		sum += float64(salary[i])
	}
	sum /= float64(len(salary) - 2)
	return sum
}

