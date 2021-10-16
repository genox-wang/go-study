package _82_expression_add_operators

import (
	"reflect"
	"testing"
)

func Test_addOperators(t *testing.T) {
	type args struct {
		num    string
		target int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test case 1",
			args{
				"123",
				6,
			},
			[]string{
				"1+2+3",
				"1*2*3",
			},
		}, {
			"test case 2",
			args{
				"232",
				8,
			},
			[]string{
				"2+3*2",
				"2*3+2",
			},
		}, {
			"test case 3",
			args{
				"105",
				5,
			},
			[]string{
				"1*0+5",
				"10-5",
			},
		}, {
			"test case 4",
			args{
				"3456237490",
				9191,
			},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOperators(tt.args.num, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOperators() = %v, want %v", got, tt.want)
			}
		})
	}
}
