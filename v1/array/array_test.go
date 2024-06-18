package array

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	var filterByGreater = func(elem int, arr []int) bool {
		return elem >= 10
	}

	type testCase[T Generics] struct {
		name string
		arg  struct {
			fn  func(elem T, arr []T) bool
			arr []T
		}
		expect []T
	}

	cases := []testCase[int]{
		{
			name: "should success",
			arg: struct{fn func(elem int, arr []int) bool; arr []int}{
				fn: filterByGreater,
				arr: []int{10, 3, 13, -5},
			},
			expect: []int{10, 13},
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

func TestFilterCurry(t *testing.T) {
	var filterByGreater = func(elem int) bool {
		return elem >= 10
	}

	type testCase[T Generics] struct {
		name string
		arg  struct {
			fn  func(elem T) bool
			arr []T
		}
		expect []T
	}

	cases := []testCase[int]{
		{
			name: "should success",
			arg: struct{fn func(elem int) bool; arr []int}{
				fn: filterByGreater,
				arr: []int{10, 3, 13, -5},
			},
			expect: []int{10, 13},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			filter := FilterCurry(c.arg.fn)

			actual := filter(c.arg.arr)
			if !reflect.DeepEqual(c.expect, actual) {
				t.Errorf("actual: %v", actual)
				t.Errorf("expect: %v", c.expect)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type testCase[T any] struct {
		name string
		arg  struct {
			fn  func(v T) T
			arr []T
		}
		expect []T
	}
	cases := []testCase[int]{
		{
			name: "should success",
			arg: struct {
				fn  func(v int) int
				arr []int
			}{
				arr: []int{1, 2, 3, 4, 5},
				fn:  func(v int) int { return v * 2 },
			},
			expect: []int{2, 4, 6, 8, 10},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Map(c.arg.fn, c.arg.arr)

			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("expect: %v", c.expect)
				t.Errorf("actual: %v", actual)
			}
		})
	}
}

func TestMapCurry(t *testing.T) {
	type testCase[T any] struct {
		name string
		arg  struct {
			fn  func(v T) T
			arr []T
		}
		expect []T
	}
	cases := []testCase[int]{
		{
			name: "should success",
			arg: struct {
				fn  func(v int) int
				arr []int
			}{
				arr: []int{1, 2, 3, 4, 5},
				fn:  func(v int) int { return v * 2 },
			},
			expect: []int{2, 4, 6, 8, 10},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			curriedMap := MapCurry(c.arg.fn)

			actual := curriedMap(c.arg.arr)

			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("expect: %v", c.expect)
				t.Errorf("actual: %v", actual)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type testCase[T any] struct {
		name string
		arg  struct {
			fn      func(acc, cur T) T
			arr     []T
			initial T
		}
		expect T
	}

	cases := []testCase[int]{
		{
			name: "should success",
			arg: struct {
				fn      func(acc int, cur int) int
				arr     []int
				initial int
			}{
				fn:      func(acc, cur int) int { return acc + cur },
				arr:     []int{1, 2, 3, 4},
				initial: 0,
			},
			expect: 10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Reduce(c.arg.fn, c.arg.arr, c.arg.initial)

			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("actual: %v", actual)
				t.Errorf("expect: %v", c.expect)
			}
		})
	}
}

func TestReduceCurry(t *testing.T) {
	type testCase[T Generics] struct {
		name string
		arg struct{
			fn func(acc, cur T) T 
			arr []T
			initial T
		}
		expect T
	}

	cases := []testCase[int]{
		{
			name: "should success",
			arg: struct{fn func(acc int, cur int) int; arr []int; initial int}{
				fn: func(acc, cur int) int {return acc+cur},
				arr: []int{1,2,3,4},
				initial: 0,
			},
			expect: 10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			reduce := ReduceCurry(c.arg.fn)

			actual := reduce(c.arg.arr, c.arg.initial)

			if !reflect.DeepEqual(actual, c.expect) {
				t.Errorf("actual: %v", actual)
				t.Errorf("expect: %v", c.expect)
			}
		})
	}
}
