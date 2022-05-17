package verr

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func createError[T Type](message string, cause error, ctx Context) *Error[T] {
	return &Error[T]{message: message, cause: cause, ctx: ctx}
}

func TestError_Error(t *testing.T) {
	dummyMessage := "uh-oh"
	type fields struct {
		message string
		cause   error
	}
	tests := []struct {
		name   string
		fields fields
		expect string
	}{
		{
			name: "When error has no cause should return expected",
			fields: fields{
				message: dummyMessage,
			},
			expect: strings.ToLower(fmt.Sprintf("[%s] %s", getTypeName(Unknown{}), dummyMessage)),
		},
		{
			name: "When error has cause, should return expected",
			fields: fields{
				message: dummyMessage,
				cause:   errors.New("cause"),
			},
			expect: strings.ToLower(fmt.Sprintf("[%s] %s: %s", getTypeName(Unknown{}), dummyMessage, "cause")),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// GIVEN
				err := createError[Unknown](tt.fields.message, tt.fields.cause, nil)

				// WHEN
				actual := err.Error()

				// THEN
				assert.Equal(t, tt.expect, actual)
			},
		)
	}
}

func TestError_Is(t *testing.T) {
	dummyError := createError[Unknown]("here's an error message", nil, nil)
	type args struct {
		err error
	}
	tests := []struct {
		name string
		err  *Error[Unknown]
		args args
		want bool
	}{
		{
			name: "When target is the same type as err, should return true",
			err:  dummyError,
			args: args{err: fmt.Errorf("uh-oh --> %w", New[Unknown]("oh no"))},
			want: true,
		},
		{
			name: "When target is not the same as err, should return false",
			err:  dummyError,
			args: args{err: errors.New("here's an error")},
			want: false,
		},
		{
			name: "When target is nil, should return false",
			err:  dummyError,
			args: args{err: nil},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equalf(t, tt.want, tt.err.Is(tt.args.err), "Is(%v)", tt.args.err)
			},
		)
	}
}
