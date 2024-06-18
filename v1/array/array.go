package array

type Generics any

func Filter[T Generics](fn func(elem T, arr []T) bool, arr []T) []T {
	if len(arr) == 0 {
		return []T{}
	}

	head := arr[0]

	next := Filter(fn, arr[1:])

	if fn(head, arr) {
		newArr := []T{head}
		newArr = append(newArr, next...)
		return newArr
	}

	return next
}

func FilterCurry[T Generics](fn func(elem T) bool) func(arr []T) []T {
	return func(arr []T) []T {
		var res []T

		for _, v := range arr {
			if fn(v) {
				res = append(res, v)
			}
		}

		return res
	}
}

func Map[T Generics](fn func(v T) T, arr []T) []T {
	if len(arr) == 0 {
		return []T{}
	}

	head := fn(arr[0])

	next := Map(fn, arr[1:])

	newArr := []T{head}
	newArr = append(newArr, next...)

	return newArr
}

func MapCurry[T Generics](fn func(v T) T) func(arr []T) []T {
	return func(arr []T) []T {
		var res []T

		for _, v := range arr {
			res = append(res, fn(v))
		}

		return res
	}
}

func Reduce[T Generics](fn func(acc, cur T) T, arr []T, initial T) T {
	res := initial
	for _, v := range arr {
		res = fn(res, v)
	}

	return res
}

func ReduceCurry[T Generics](fn func(acc, cur T) T) func(arr []T, initial T) T {
	return func(arr []T, initial T) T {
		res := initial

		for _, v := range arr {
			res = fn(res, v)
		}

		return res
	}
}
