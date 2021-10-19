package gomod

import "testing"

func Test_parseGoMod(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"test case 1",
			"1.15",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseGoMod(); got != tt.want {
				t.Errorf("parseGoMod() = %v, want %v", got, tt.want)
			}
		})
	}
}
