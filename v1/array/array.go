package array

func Filter(fn func(elem any, arr []any) bool, arr []any) []any {
	if len(arr) == 0 {
		return []any{}
	}

	head := arr[0]

	next := Filter(fn, arr[1:])

	if fn(head, arr) {
		newArr := []any{head}
		newArr = append(newArr, next...)
		return newArr
	}

	return next
}
