package lemoine

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	if IsPrime(97) != true {
		t.Error("97 fail")
	}
	if IsPrime(98) != false {
		t.Error("98 fail")
	}
}

func TestLemoine(t *testing.T) {
	Lemoine(39)
	Lemoine(7)
	Lemoine(101)
}
