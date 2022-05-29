package leetcode

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"", "()()", true},
		{"", "([[)", false},
		{"", "([{}])", true},
		{"", "((((", false},
		{"", "{}{}(){}", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
