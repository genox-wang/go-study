package template

import (
	"testing"
)

func TestWhiteCat(t *testing.T) {
	w := &WhiteCat{}
	if w.eat() != "Cat eat" {
		t.Errorf("111")
	}
	if w.run() != "WhiteCat run" {
		t.Errorf("222")
	}
	if w.jump() != "Cat jump" {
		t.Error("333")
	}
}
