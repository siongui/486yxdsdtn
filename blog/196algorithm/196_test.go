package lychrel

import (
	"testing"
)

func TestReverse(t *testing.T) {
	if Reverse("你好嗎") != "嗎好你" {
		t.Error("你好嗎")
		return
	}
}

func TestLychrelNumberTest(t *testing.T) {
	LychrelNumberTest(56, 20)
	LychrelNumberTest(87, 20)
}
