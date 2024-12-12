package util

func New[T any](e T) *T {
	return &e
}

func Must[T any](t T, e any) T {
	if e != nil {
		switch t := e.(type) {
		case bool:
			if !t {
				panic(t)
			}
		default:
			panic(e)
		}
	}
	return t
}

// Combinations returns combinations of n elements for a given slice.
// For n < 1, returns an empty slice.
// For n > len(slice), it is equal to n = len(slice).
func Combinations[T any](slice []T, n int) (rt [][]T) {
	pool := slice
	l := len(pool)

	if n > l {
		n = l
	}

	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	result := make([]T, n)
	for i, el := range indices {
		result[i] = pool[el]
	}
	s2 := make([]T, n)
	copy(s2, result)
	rt = append(rt, s2)

	for {
		i := n - 1
		for ; i >= 0 && indices[i] == i+l-n; i -= 1 {
		}

		if i < 0 {
			return
		}

		indices[i] += 1
		for j := i + 1; j < n; j += 1 {
			indices[j] = indices[j-1] + 1
		}

		for ; i < len(indices); i += 1 {
			result[i] = pool[indices[i]]
		}
		s2 = make([]T, n)
		copy(s2, result)
		rt = append(rt, s2)
	}
}
