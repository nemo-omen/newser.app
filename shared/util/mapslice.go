package util

func MapSlice[T any, M any](arr []T, mapFn func(T) M) []M {
	newSlice := make([]M, len(arr))
	for i, elem := range arr {
		newSlice[i] = mapFn(elem)
	}
	return newSlice
}
