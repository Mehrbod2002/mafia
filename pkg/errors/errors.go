package errors

import "errors"

var (
	// ErrNotFound represents missing resources.
	ErrNotFound = errors.New("not found")
	// ErrUnauthorized signals missing or invalid credentials.
	ErrUnauthorized = errors.New("unauthorized")
	// ErrForbidden indicates the caller lacks permissions.
	ErrForbidden = errors.New("forbidden")
	// ErrConflict communicates conflicting state like duplicates.
	ErrConflict = errors.New("conflict")
	// ErrInvalid is returned when input validation fails.
	ErrInvalid = errors.New("invalid")
)
