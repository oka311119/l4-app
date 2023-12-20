package auth

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrFailedSaltGeneration = errors.New("failed to generate salt")
)
