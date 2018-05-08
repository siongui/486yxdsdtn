package diff

import (
	"testing"
)

func TestDiff(t *testing.T) {
	before := []string{"hello", "abc", "world"}
	after := []string{"hello", "world"}

	pairs := Diff(before, after)
	for _, pair := range pairs {
		t.Log(pair)
	}
}
