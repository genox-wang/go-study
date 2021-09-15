package singleton

import (
	"reflect"
	"testing"
)

func TestGetLazyOnceSingletonOr(t *testing.T) {
	tests := []struct {
		name string
		want *LazyOnceSingleton
	}{
		{
			name: "LazyOnceSingleton test",
			want: GetLazyOnceSingletonOr(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLazyOnceSingletonOr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLazyOnceSingletonOr() = %v, want %v", got, tt.want)
			}
		})
	}
}
