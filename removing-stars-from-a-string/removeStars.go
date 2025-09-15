/*
	Package t2390

* You are given a string s, which contains stars *.

In one operation, you can:

Choose a star in s.
Remove the closest non-star character to its left, as well as remove the star itself.
Return the string after all stars have been removed.

Note:

The input will be generated such that the operation is always possible.
It can be shown that the resulting string will always be unique.
*/
package t2390

func removeStars(s string) string {
	chars := make([]byte, len(s))
	write := 0

	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			chars[write] = s[i]
			write++
		} else {
			write--
		}

		if write < 0 {
			return ""
		}
	}

	return string(chars[:write])
}
