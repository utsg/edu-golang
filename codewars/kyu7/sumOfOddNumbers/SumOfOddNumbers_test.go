package kyu7

import "testing"

func TestRowSumOddNumbers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args int
		want int
	}{
		{"", 1, 1},
		{"", 2, 8},
		{"", 13, 2197},
		{"", 19, 6859},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RowSumOddNumbers(tt.args); got != tt.want {
				t.Errorf("RowSumOddNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
