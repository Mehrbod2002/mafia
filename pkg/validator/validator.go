package validator

import "regexp"

var emailRegex = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

// Email validates a basic email format.
func Email(value string) bool {
	return emailRegex.MatchString(value)
}

// NonEmpty ensures the string contains characters.
func NonEmpty(value string) bool {
	return len(value) > 0
}

// MinLen verifies a string meets a minimum length.
func MinLen(value string, length int) bool {
	return len(value) >= length
}
