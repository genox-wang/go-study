package strategy

import (
	"testing"
)

func TestBehavior(t *testing.T) {
	strategy := &Strategy{}
	strategy.setBehavior(&Run{})
	got := strategy.Do()
	want := "Run"
	if got != want {
		t.Errorf("Strategy.Do() = %v, want %v", got, want)
	}
	strategy.setBehavior(&Jump{})
	got = strategy.Do()
	want = "Jump"
	if got != want {
		t.Errorf("Strategy.Do() = %v, want %v", got, want)
	}
}
