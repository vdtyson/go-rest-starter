package verr

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithContext(t *testing.T) {
	type args struct {
		ctx     Context
		message string
		args    []any
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "When context provided, should return expected",
			args: args{
				ctx:     Context{"a": "b"},
				message: "Uh-oh %s",
				args:    []any{"what's going on?"},
			},
			want: &Error[Unknown]{errType: Unknown{}, message: "Uh-oh what's going on?", ctx: Context{"a": "b"}},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// GIVEN
				expected := tt.want

				// WHEN
				actual := WithContext[Unknown](tt.args.ctx, tt.args.message, tt.args.args...)

				// THEN
				assert.Equal(t, expected, actual)
			},
		)
	}
}

func TestNew(t *testing.T) {
	type args struct {
		message string
		args    []interface{}
	}
	tests := []struct {
		name   string
		args   args
		expect error
	}{
		{
			name: "Should return expected",
			args: args{
				message: "uh-oh %s",
				args:    []any{"what happened?"},
			},
			expect: createError[Unknown]("uh-oh what happened?", nil, nil),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// WHEN
				actual := New[Unknown](tt.args.message, tt.args.args...)

				// THEN
				assert.Equal(t, tt.expect, actual)
			},
		)
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		cause error
		msg   string
		args  []interface{}
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "When cause is nil should return nil",
			args: args{cause: nil},
			want: nil,
		},
		{
			name: "When cause exists should return expected",
			args: args{cause: fmt.Errorf("dkdk"), msg: "abc"},
			want: newError[Unknown](nil, fmt.Errorf("dkdk"), "abc"),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// WHEN
				actual := Wrap(tt.args.cause, tt.args.msg, tt.args.args...)

				// THEN
				assert.Equal(t, tt.want, actual)
			},
		)
	}
}
