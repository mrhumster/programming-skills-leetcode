/*
Package t50

Implement pow(x, n), which calculates x raised to the power n (i.e., xn).
*/
package t50

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 == 1 {
		return x * myPow(x*x, n/2)
	}
	return myPow(x*x, n/2)
}
