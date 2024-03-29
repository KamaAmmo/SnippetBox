package models

import "errors"

var (
	ErrNoRecord error = errors.New("models: no matching record found")

	ErrInvalidCredentials error = errors.New("models: invalid credentials")

	ErrDuplicateEmail error = errors.New("models: duplicate emails")
)
