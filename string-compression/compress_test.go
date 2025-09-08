package t443

import "testing"

func TestCompress(t *testing.T) {
	data := []struct {
		name     string
		chars    []byte
		expected int
	}{
		{"1", []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6},
		{"2", []byte{'a'}, 1},
		{"3", []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, 4},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := compress(d.chars)
			if resp != d.expected {
				t.Errorf(
					"⚠️ Test `%s` failed. Expected `%d`, but got `%d`.",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}

