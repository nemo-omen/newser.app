package util

func Filter[T any](s []T, test func(t T) bool) (ret []T) {
	for _, el := range s {
		if test(el) {
			ret = append(ret, el)
		}
	}
	return
}
