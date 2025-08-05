/*
Package t66
You are given a large integer represented as an integer array digits, where
Paeach digits[i] is the ith digit of the integer. The digits are ordered
from most significant to least significant in left-to-right order. The large
integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.
*/
package t66

import "fmt"

func moreThatTen(i int) (bool, int, int) {
	more := false
	ten := 0
	unit := i
	if i >= 10 {
		more = true
		ten = i / 10
		unit = i % 10
	}
	return more, ten, unit
}

func shiftSlice(a *[]int) {
	fmt.Println("Before:", *a)
	*a = append(*a, 0)
	for i := len(*a) - 1; i > 0; i-- {
		(*a)[i] = (*a)[i-1]
		(*a)[i-1] = 0
	}
	fmt.Println("After:", *a)
}

func checkingForOverflow(a *[]int) {
	l := len(*a)
	for i := l - 1; i >= 0; i-- {
		more, ten, unit := moreThatTen((*a)[i])
		if !more {
			continue
		}
		if i == 0 {
			shiftSlice(a)
			(*a)[i+1] = unit
			(*a)[i] += ten
		} else {
			(*a)[i] = unit
			(*a)[i-1] += ten
		}
	}
}

func plusOne(digits []int) []int {
	l := len(digits)
	digits[l-1] += 1
	checkingForOverflow(&digits)
	return digits
}
