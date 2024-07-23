package aliU

func MergeMaps[T comparable, U any](a map[T]U, b map[T]U) map[T]U {
	if a == nil {
		a = map[T]U{}
	}
	for k, v := range b {
		a[k] = v
	}
	return a
}
