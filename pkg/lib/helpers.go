package lib

import "math/rand"

// Coalesce returns the first value if it is not the zero value for it's type,
// otherwise returns the second value.
func Coalesce[T comparable](v1, v2 T) T {
	var zero T
	if v1 != zero {
		return v1
	}
	return v2
}

// IIF is an inline if statement that returns the first value if the condition
// is true, otherwise returns the second value.
func IIF[T any](condition bool, v1, v2 T) T {
	if condition {
		return v1
	}
	return v2
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GenerateUniqueString returns a random string of upper and lower case English
// letters of the given length.
func GenerateUniqueString(length uint) string {
	s := make([]rune, int(length))

	for i := 0; i < int(length); i++ {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}
