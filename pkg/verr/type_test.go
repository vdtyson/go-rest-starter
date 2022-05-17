package verr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getTypeName(t *testing.T) {
	type ErrTest struct {
		Type
	}
	type args struct {
		t Type
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "When passed an ErrTest Type, should return expected",
			args: args{t: ErrTest{}},
			want: "ErrTest",
		},
		{
			name: "When passed pointer to ErrTest, should return expected",
			args: args{t: &ErrTest{}},
			want: "ErrTest",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := getTypeName(tt.args.t); got != tt.want {
					t.Errorf("getTypeName() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_typesEqual(t *testing.T) {
	type ErrType1 struct {
		Type
	}
	type ErrType2 struct {
		Type
	}
	type args struct {
		t1 Type
		t2 Type
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "When equal, should return true",
			args: args{t1: ErrType1{}, t2: ErrType1{}},
			want: true,
		},
		{
			name: "When not equal, should return false",
			args: args{t1: ErrType1{}, t2: ErrType2{}},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equalf(
					t,
					tt.want,
					typesEqual(tt.args.t1, tt.args.t2),
					"typesEqual(%v, %v)",
					tt.args.t1,
					tt.args.t2,
				)
			},
		)
	}
}
