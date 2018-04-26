package db3

import (
	"testing"
)

func TestIsIntDivisibleBy3(t *testing.T) {
	if IsIntDivisibleBy3(123456758933312) != false {
		t.Error(123456758933312)
	}
	if IsIntDivisibleBy3(769452) != true {
		t.Error(769452)
	}
}

func TestIsStrDivisibleBy3(t *testing.T) {
	result, err := IsStrDivisibleBy3("3635883959606670431112222")
	if err != nil {
		t.Error(err)
	}
	if result != true {
		t.Error("3635883959606670431112222")
	}

	result, err = IsStrDivisibleBy3("123456758933312")
	if err != nil {
		t.Error(err)
	}
	if result != false {
		t.Error("123456758933312")
	}

	result, err = IsStrDivisibleBy3("769452")
	if err != nil {
		t.Error(err)
	}
	if result != true {
		t.Error("769452")
	}
}
