package t1491

import (
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {
	data := []struct {
		name     string
		salary   []int
		expected float64
	}{
		{
			"Ex.1",
			[]int{4000, 3000, 1000, 2000},
			2500.00000,
		},
		{
			"Ex.2",
			[]int{1000, 2000, 3000},
			2000.00000,
		},
	}
	for _, d := range data {
		resp := average(d.salary)
		if !reflect.DeepEqual(d.expected, resp) {
			t.Errorf(
				"\n⚠️ Test `%s` failed! Expected `%f`, but got `%f`.",
				d.name,
				d.expected,
				resp,
			)
		}
	}
}

