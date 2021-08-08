// Package app use cases
package app

import "errors"

var (
	// ErrInvalid input.
	ErrInvalid = errors.New("invalid")
	// ErrOnSave situation.
	ErrOnSave = errors.New("on save")
	// ErrOnGenerateToken auth token.
	ErrOnGenerateToken = errors.New("on generate token")
	// ErrOnSendMessage to user.
	ErrOnSendMessage = errors.New("on send message")
)
