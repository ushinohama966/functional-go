package array

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	var filterByGreater = func(elem any, arr []any)bool{
		n, ok := elem.(int)
		if !ok  {
			return false
		}

		return n >= 10
	}

	cases := []struct {
		name string
		arg  struct {
			fn  func(elem any, arr []any) bool
			arr []any
		}
		expect []any
	}{
		{
			name: "",
			arg: struct {
				fn  func(elem any, arr []any) bool
				arr []any
			}{
				arr: []any{10, 13, 5, -9},
				fn: filterByGreater,
			},
			expect: []any{10,13},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Filter(c.arg.fn, c.arg.arr)
			if !reflect.DeepEqual(c.expect, actual) {
				t.Errorf("actual: %v", actual)
				t.Errorf("expect: %v", c.expect)
			}
		})
	}
}
