/*
Package t1431

There are n kids with candies. You are given an integer array candies,
where each candies[i] represents the number of candies the ith kid has,
and an integer extraCandies, denoting the number of extra candies that you have.

Return a boolean array result of length n, where result[i] is true if, after
giving the ith kid all the extraCandies, they will have the greatest number
of candies among all the kids, or false otherwise.

Note that multiple kids can have the greatest number of candies.
*/
package t1431

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	mx := slices.Max(candies)
	var response []bool
	for _, v := range candies {
		response = append(response, v+extraCandies >= mx)
	}
	return response
}
