package factory

import (
	"testing"
)

func Test_dioerDo(t *testing.T) {
	type args struct {
		dioer Dioer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"dioer HttpClient test",
			args{
				dioer: NewHttpClient(),
			},
			"HttpClient",
		},
		{
			"dioer MockHttpClient test",
			args{
				dioer: NewMockHttpClient(),
			},
			"MockHttpClient",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dioerDo(tt.args.dioer); got != tt.want {
				t.Errorf("dioerDo() = %v, want %v", got, tt.want)
			}
		})
	}
}
