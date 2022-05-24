package kyu7

import "testing"

func TestPrinterError(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz", "3/56"},
		{"", "kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz", "6/60"},
		{"", "kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu", "11/65"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrinterError(tt.args); got != tt.want {
				t.Errorf("PrinterError() = %v, want %v", got, tt.want)
			}
		})
	}
}
