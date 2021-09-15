package option

import (
	"reflect"
	"testing"
)

func TestNewConnection(t *testing.T) {
	type args struct {
		addr string
		os   []Option
	}
	tests := []struct {
		name string
		args args
		want *Connection
	}{
		{
			"TestNewConnection",
			args{
				addr: "localhost",
				os: []Option{
					WithCache(false),
					WithTimeout(1010),
				},
			},
			&Connection{
				addr:    "localhost",
				cache:   false,
				timeout: 1010,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConnection(tt.args.addr, tt.args.os...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
