/*
Package t43
Given two non-negative integers num1 and num2 represented as strings,
return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/
package t43

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func multiplySimple(a, b string) string {
	if len(a) > 1 || len(b) > 1 {
		fmt.Println("For digits only")
		os.Exit(1)
	}
	ad, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bd, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strconv.Itoa(ad * bd)
}

func sumSimple(a, b string) string {
	ad, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bd, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strconv.Itoa(ad + bd)
}

func splitNumber(a string) (string, string) {
	if len(a) == 1 {
		return "0", a
	}
	if len(a) > 2 {
		fmt.Println("Applicable only for 0-99")
		os.Exit(1)
	}
	return string(a[0]), string(a[1])
}

func amount(a, b string) (string, error) {
	if len(b) > len(a) {
		a, b = b, a
	}
	b = strings.Repeat("0", len(a)-len(b)) + b
	// fmt.Println("a=", a)
	// fmt.Println("b=", b)
	carry := "0"
	totalSum := ""
	for i := len(a) - 1; i >= 0; i-- {
		sum := sumSimple(string(a[i]), string(b[i]))
		sum = sumSimple(sum, carry)
		tens, digits := splitNumber(sum)
		// fmt.Printf("a: %s, b: %s, carry: %s, sum: %s, totalSum: %s\n", a, b, carry, sum, totalSum)
		totalSum = digits + totalSum
		carry = tens
	}
	if carry != "0" {
		totalSum = carry + totalSum
	}
	return totalSum, nil
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	lines := []string{}
	if len(num2) > len(num1) {
		num1, num2 = num2, num1
	}
	level := 0
	for i := len(num1) - 1; i >= 0; i-- {
		t := "0"
		lines = append(lines, "")
		for j := len(num2) - 1; j >= 0; j-- {
			a := string(num1[i])
			b := string(num2[j])
			product := multiplySimple(a, b)
			product = sumSimple(product, t)
			fmt.Printf("%s * %s = %s\n", a, b, product)
			tens, digits := splitNumber(product)
			t = tens
			lines[level] = digits + lines[level]
			if j == 0 && tens != "0" {
				lines[level] = tens + lines[level]
			}
			if j == 0 {
				lines[level] = lines[level] + strings.Repeat("0", level)
			}
		}
		level += 1
	}
	// fmt.Printf("%#v\n", lines)
	// ---- PLUS
	finalSum := "0"
	for _, line := range lines {
		// fmt.Printf("%d %s + %s\n", i, finalSum, line)
		sum, err := amount(finalSum, line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		finalSum = sum
	}
	return finalSum
}
