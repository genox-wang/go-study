package _48_shortest_completing_word

import "testing"

func Test_shortestCompletingWord(t *testing.T) {
	type args struct {
		licensePlate string
		words        []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test case 1",
			args{
				"1s3 PSt",
				[]string{"step", "steps", "stripe", "stepple"},
			},
			"steps",
		}, {
			"test case 2",
			args{
				"1s3 456",
				[]string{"looks", "pest", "stew", "show"},
			},
			"pest",
		}, {
			"test case 3",
			args{
				"OgEu755",
				[]string{"enough", "these", "play", "wide", "wonder", "box", "arrive", "money", "tax", "thus"},
			},
			"enough",
		}, {
			"test case 4",
			args{
				"iMSlpe4",
				[]string{"claim", "consumer", "student", "camera", "public", "never", "wonder", "simple", "thought", "use"},
			},
			"simple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestCompletingWord(tt.args.licensePlate, tt.args.words); got != tt.want {
				t.Errorf("shortestCompletingWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
