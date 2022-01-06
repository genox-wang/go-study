package _1_simplify_path

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test case 1",
			args{"/home/"},
			"/home",
		}, {
			"test case 2",
			args{"/../"},
			"/",
		}, {
			"test case 3",
			args{"/home//foo/"},
			"/home/foo",
		}, {
			"test case 4",
			args{"/a/./b/../../c/"},
			"/c",
		}, {
			"test case 5",
			args{"/a//b////c/d//././/.."},
			"/a/b/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
