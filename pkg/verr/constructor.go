package verr

import (
	"fmt"
)

func New[T Type](message string, args ...any) error {
	return newError[T](nil, nil, message, args...)
}

func WithContext[T Type](ctx Context, message string, args ...any) error {
	return newError[T](ctx, nil, message, args...)
}

func Wrap(cause error, msg string, args ...any) error {
	if cause == nil {
		return nil
	}
	return WrapT[Unknown](cause, msg, args...)
}

func WrapT[T Type](cause error, msg string, args ...any) error {
	if cause == nil {
		return nil
	}
	return newError[T](nil, cause, msg, args...)
}

func newError[T Type](ctx Context, cause error, message string, args ...any) *Error[T] {
	var t T
	return &Error[T]{errType: t, ctx: ctx, cause: cause, message: fmt.Sprintf(message, args...)}
}
