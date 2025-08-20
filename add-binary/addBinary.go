/*
	Package t67

Given two binary strings a and b, return their sum as a binary string.
*/
package t67

import (
	"strings"
)

func sumSimple(a, b string) (string, string) {
	if a == "0" && b == "0" {
		return "0", "0"
	}
	if a == "0" && b == "1" {
		return "0", "1"
	}
	if a == "1" && b == "0" {
		return "0", "1"
	}
	if a == "1" && b == "1" {
		return "1", "0"
	}
	return "-1", "-1"
}

func sumThree(a, b, carry string) (string, string) {
	carry1, sum1 := sumSimple(a, b)
	carry2, sum2 := sumSimple(sum1, carry)
	_, finalCarry := sumSimple(carry1, carry2)
	return finalCarry, sum2
}

func reverseWithBuilder(s string) string {
	var sb strings.Builder
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	if lb > la {
		a = strings.Repeat("0", lb-la) + a
	}
	if la > lb {
		b = strings.Repeat("0", la-lb) + b
	}
	carry := "0"
	var result strings.Builder

	for i := len(a) - 1; i >= 0; i-- {

		carry1, sum := sumThree(string(a[i]), string(b[i]), carry)
		result.WriteString(sum)
		carry = carry1
	}
	if carry == "1" {
		result.WriteString("1")
	}
	return reverseWithBuilder(result.String())
}
