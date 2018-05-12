package politeness

import (
	"testing"
)

func TestCalculatePoliteness(t *testing.T) {
	if CalculatePoliteness(9) != 2 {
		t.Error("politeness of 9 is not 2")
		return
	}
	if CalculatePoliteness(15) != 3 {
		t.Error("politeness of 15 is not 3")
		return
	}
	if CalculatePoliteness(90) != 5 {
		t.Error("politeness of 90 is not 5")
		return
	}
}
