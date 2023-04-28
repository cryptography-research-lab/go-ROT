package rot

import "testing"

func TestROTForNumber(t *testing.T) {
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
				value: '0',
				rotN:  1,
			},
			want: '1',
		},
		testCase[byte]{
			args: args[byte]{
				value: '9',
				rotN:  1,
			},
			want: '0',
		},
		testCase[byte]{
			args: args[byte]{
				value: 'A',
				rotN:  1,
			},
			want: 'A',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ROTForNumber(tt.args.value, tt.args.rotN); got != tt.want {
				t.Errorf("ROTForNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
