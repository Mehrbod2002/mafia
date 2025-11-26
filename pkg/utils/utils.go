package utils

// Ptr returns a pointer to the provided value.
func Ptr[T any](v T) *T {
	return &v
}

// Default returns value if not zero-value else fallback.
func Default[T comparable](value T, fallback T) T {
	var zero T
	if value == zero {
		return fallback
	}
	return value
}
