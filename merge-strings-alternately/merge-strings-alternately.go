package T1768

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	if len(word1) > 100 || len(word2) > 100 || word1 == "" || word2 == "" {
		fmt.Println("⚠️ Error! The parametr does not pass uunder the condition.")
		return ""
	}
	var _l int
	var _mixed string = ""
	l1 := len(word1)
	l2 := len(word2)
	if l1 >= l2 {
		_l = l1
	} else {
		_l = l2
	}
	for i := 0; i <= _l; i++ {
		if i < l1 {
			_mixed = _mixed + string(word1[i])
		}
		if i < l2 {
			_mixed = _mixed + string(word2[i])
		}
	}
	return _mixed
}
