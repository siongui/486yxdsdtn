package getpasswd

import (
	"testing"
)

func TestGetpasswd(t *testing.T) {
	pwd, err := Getpasswd()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(pwd)
}
