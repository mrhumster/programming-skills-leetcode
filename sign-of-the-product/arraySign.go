/*
Package t1822
1822. Sign of the Product of an Array
Implement a function signFunc(x) that returns:

1 if x is positive.
-1 if x is negative.
0 if x is equal to 0.
You are given an integer array nums. Let product be the product of all values in the array nums.

Return signFunc(product).
*/
package t1822

func signFunc(a *[]int) int {
	p := 1

	for _, v := range *a {
		if v == 0 {
			p = 0
			break
		}
		if v < 0 {
			p *= -1
		}
	}

	if p > 0 {
		return 1
	}
	if p < 0 {
		return -1
	}
	return 0
}

func arraySign(nums []int) int {
	return signFunc(&nums)
}
