package util

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Power[T Number, P int](n T, p P) T {
	if p == 0 {
		return 1
	}
	return n * Power[T](n, p-1)
}

func Abs[T Number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func Digits[T Number](n T) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}
