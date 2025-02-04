package foo

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 != 3")
	}
}

func TestAddFail(t *testing.T) {
	if Add(1, 2) != 4 {
		t.Error("1 + 2 != 4")
	}
}
