package t58

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	data := []struct {
		name     string
		s        string
		expected int
	}{
		{
			"Ex.1", "Hello World", 5,
		},
		{
			"Ex.2", "   fly me   to   the moon  ", 4,
		},
		{
			"Ex.3", "luffy is still joyboy", 6,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := lengthOfLastWord(d.s)
			if resp != d.expected {
				t.Errorf("⚠️ Test `%s` failed! Expected %d, but got %d.", d.name, d.expected, resp)
			}
		})
	}
}
