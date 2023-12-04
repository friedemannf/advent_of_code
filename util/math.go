package util

func Power[T ~int](n, p T) T {
	if p == 0 {
		return 1
	}
	return n * Power(n, p-1)
}
