package rot

import (
	"testing"
)

func TestROTForLetter(t *testing.T) {
	type args[T interface{ byte | rune }] struct {
		value T
		rotN  int
	}
	type testCase[T interface{ byte | rune }] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[byte]{
		testCase[byte]{
			args: args[byte]{
				value: 'a',
				rotN:  1,
			},
			want: 'b',
		},
		testCase[byte]{
			args: args[byte]{
				value: 'z',
				rotN:  1,
			},
			want: 'a',
		},
		testCase[byte]{
			args: args[byte]{
				value: 'A',
				rotN:  1,
			},
			want: 'B',
		},
		testCase[byte]{
			args: args[byte]{
				value: 'Z',
				rotN:  1,
			},
			want: 'A',
		},
		testCase[byte]{
			args: args[byte]{
				value: '1',
				rotN:  1,
			},
			want: '1',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ROTForLetter(tt.args.value, tt.args.rotN); got != tt.want {
				t.Errorf("ROTForLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
