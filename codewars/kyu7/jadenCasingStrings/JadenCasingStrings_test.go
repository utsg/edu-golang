package kyu7

import (
	"testing"
)

func ToJadenCaseTest(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		str  string
		want string
	}{
		{"", "most trees are blue", "Most Trees Are Blue"},
		{"", "All the rules in this world were made by someone no smarter than you. So make your own.", "All The Rules In This World Were Made By Someone No Smarter Than You. So Make Your Own."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToJadenCase(tt.str); got != tt.want {
				t.Errorf("ToJadenCase(): [%v], want: [%v]", got, tt.want)
			}
		})
	}
}
