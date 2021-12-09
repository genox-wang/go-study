package _94_valid_tic_tac_toe_state

import "testing"

func Test_validTicTacToe(t *testing.T) {
	type args struct {
		board []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test case 1",
			args{
				[]string{"O  ", "   ", "   "},
			},
			false,
		}, {
			"test case 2",
			args{
				[]string{"XOX", " X ", "   "},
			},
			false,
		}, {
			"test case 3",
			args{
				[]string{"XXX", "   ", "OOO"},
			},
			false,
		}, {
			"test case 4",
			args{
				[]string{"XOX", "O O", "XOX"},
			},
			true,
		}, {
			"test case 5",
			args{
				[]string{"XXX", "OOX", "OOX"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validTicTacToe(tt.args.board); got != tt.want {
				t.Errorf("validTicTacToe() = %v, want %v", got, tt.want)
			}
		})
	}
}
