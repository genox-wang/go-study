package singleton

import (
	"reflect"
	"testing"
)

func TestGetLazySingletonOr(t *testing.T) {
	tests := []struct {
		name string
		want *LazySingleton
	}{
		{
			name: "LazySingleton test",
			want: GetLazySingletonOr(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLazySingletonOr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLazySingletonOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
