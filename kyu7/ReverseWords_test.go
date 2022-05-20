package kyu7

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		str  string
		want string
	}{
		{"", "The quick brown fox jumps over the lazy dog.", "ehT kciuq nworb xof spmuj revo eht yzal .god"},
		{"", "apple", "elppa"},
		{"", "a b c d", "a b c d"},
		{"", "double  spaced  words", "elbuod  decaps  sdrow"},
		{"", "stressed desserts", "desserts stressed"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseWords(tt.str); got != tt.want {
				t.Errorf("ReverseWords(): [%v], want: [%v]", got, tt.want)
			}
		})
	}
}
