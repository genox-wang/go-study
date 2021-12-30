package _46_hand_of_straights

import "testing"

func Test_isNStraightHand(t *testing.T) {
	type args struct {
		hand      []int
		groupSize int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test case 1",
			args{
				[]int{1, 2, 3, 6, 2, 3, 4, 7, 8},
				3,
			},
			true,
		}, {
			"test case 2",
			args{
				[]int{1, 2, 3, 4, 5},
				4,
			},
			false,
		}, {
			"test case 3",
			args{
				[]int{1, 1, 2, 2, 3, 3},
				3,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNStraightHand(tt.args.hand, tt.args.groupSize); got != tt.want {
				t.Errorf("isNStraightHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
