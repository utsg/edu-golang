package kyu7

import "testing"

func TestSolve(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args string
		want rune
	}{
		{"", "a", 97},
		{"", "aa", 97},
		{"", "bcd", 98},
		{"", "axyzxyz", 120},
		{"", "dcbadcba", 97},
		{"", "aabccc", 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args); got != tt.want {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
