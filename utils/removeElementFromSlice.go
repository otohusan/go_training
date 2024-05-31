package utils

// RemoveElement removes the element at index i from slice s.
func RemoveElementFromSlice[T any](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}
