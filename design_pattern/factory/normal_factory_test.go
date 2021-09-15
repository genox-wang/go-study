package factory

import (
	"reflect"
	"testing"
)

func TestPersion_GetName(t *testing.T) {
	tests := []struct {
		name string
		p    Persion
		want string
	}{
		{
			"normal_fatory test",
			Persion{Name: "genox"},
			"genox",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GetName(); got != tt.want {
				t.Errorf("Persion.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPersion(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Persion
	}{
		{
			"normal_factory new person test",
			args{
				name: "genox",
			},
			&Persion{Name: "genox"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPersion(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
