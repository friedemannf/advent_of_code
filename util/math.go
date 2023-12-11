package util

func Power[T ~int](n, p T) T {
	if p == 0 {
		return 1
	}
	return n * Power(n, p-1)
}

func Abs[T ~int](n T) T {
	if n < 0 {
		return -n
	}
	return n
}
