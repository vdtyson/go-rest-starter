package verr

import (
	"errors"
	"fmt"
	"strings"
)

type IError interface {
	Error() string
	Unwrap() error
	Is(e error) bool
}

// TODO: Add stack trace
type Error[T Type] struct {
	errType T
	message string
	ctx     Context
	cause   error
}

func (e *Error[T]) Error() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("[%s] %s", strings.ToLower(getTypeName(e.errType)), e.message))

	if e.cause != nil {
		builder.WriteString(fmt.Sprintf(": %s", e.cause.Error()))
	}

	return builder.String()
}

func (e *Error[T]) Unwrap() error {
	return e.cause
}

// Is reports whether any error in err's chain matches target.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
func (e *Error[T]) Is(err error) bool {
	if err == nil {
		return false
	}
	target := err
	for target != nil {
		if getElementType(e) == getElementType(target) {
			return true
		}
		target = errors.Unwrap(err)
	}
	return false
}

func (e *Error[T]) IsType(target Type) bool {
	var curr error = e
	for curr != nil {
		if IsType(e.errType, target) {
			return true
		}
		curr = e.Unwrap()
	}
	return false
}

// TODO: Implement format method
/*// Format implements fmt.Formatter interface
func (e *Error[T]) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		if st.Flag('+') {
			io.WriteString(st, e.Error())
		}
	}
}*/
