package util

func New[T any](e T) *T {
	return &e
}
