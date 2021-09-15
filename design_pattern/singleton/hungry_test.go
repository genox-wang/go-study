package singleton

import (
	"reflect"
	"testing"
)

func TestGetInstOr(t *testing.T) {
	tests := []struct {
		name string
		want *HungrySingleton
	}{
		{
			name: "HungrySingleton test",
			want: GetInstOr(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstOr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
