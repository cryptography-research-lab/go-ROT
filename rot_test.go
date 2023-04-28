package rot

import (
	"github.com/golang-infrastructure/go-gtypes"
	"math"
	"testing"
)

func TestROT(t *testing.T) {
	type args[T gtypes.Integer] struct {
		value T
		rotN  T
	}
	type testCase[T gtypes.Integer] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		testCase[int]{
			args: args[int]{
				value: 1,
				rotN:  100,
			},
			want: 101,
		},
		testCase[int]{
			args: args[int]{
				value: math.MaxInt,
				rotN:  1,
			},
			want: math.MinInt,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ROT(tt.args.value, tt.args.rotN); got != tt.want {
				t.Errorf("ROT() = %v, want %v", got, tt.want)
			}
		})
	}
}
