package util

func Map[A any, B any](fn func(A) B, values []A) []B {
	result := make([]B, 0, len(values))
	for _, v := range values {
		result = append(result, fn(v))
	}
	return result
}

func Filter[A any](fn func(A) bool, values []A) []A {
	result := make([]A, 0, len(values))
	for _, v := range values {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func Fold[A any, B any](fn func(A, B) B, b B, values []A) B {
	var result = b
	for _, a := range values {
		result = fn(a, b)
	}
	return result
}
