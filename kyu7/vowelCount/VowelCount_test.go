package main

import "testing"

func TestGetCount(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name      string
		str       string
		wantCount int
	}{
		{"", "abracadabra", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := GetCount(tt.str); gotCount != tt.wantCount {
				t.Errorf("GetCount() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
