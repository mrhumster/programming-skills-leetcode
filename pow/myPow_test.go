package t50

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	data := []struct {
		name     string
		x        float64
		n        int
		expected float64
	}{
		{
			"Ch.1",
			2.00000,
			10,
			1024.00000,
		},
		{
			"Ch.2",
			2.10000,
			3,
			9.26100,
		},
		{
			"Ch.3",
			2.00000,
			-2,
			0.25,
		},
		{
			"Ch.4",
			1.00000,
			2147483647,
			1.00000,
		},
		{
			"Ch.5",
			2.00000,
			-2147483648,
			0e-10,
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			resp := myPow(d.x, d.n)
			if math.Abs(d.expected-resp) > 1e-9 {
				t.Errorf(
					"⚠️ Test `%s` failed! Expected  `%e`, but got `%e`!",
					d.name,
					d.expected,
					resp,
				)
			}
		})
	}
}
