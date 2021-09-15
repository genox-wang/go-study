package factory

import (
	"reflect"
	"testing"
)

func TestNewStudent(t *testing.T) {
	type args struct {
		name  string
		score int
	}
	tests := []struct {
		name string
		args args
		want *Student
	}{
		{
			"TestNewStudent test",
			args{
				name:  "genox",
				score: 100,
			},
			&Student{
				Name:  "genox",
				Score: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStudent(tt.args.name)(tt.args.score); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}
