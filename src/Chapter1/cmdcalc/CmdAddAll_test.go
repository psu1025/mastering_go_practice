package cmdcalc

import "testing"

func Test_AddAllVariable(t *testing.T) {
	numstrings := []string{"1", "2", "3", "ㄱ"}
	var total, length = AddAllVariable(numstrings)

	if total != 6 {
		t.Fatal("1+2+3 == 6")
	}

	if length != 3 {
		t.Fatal("length must be 3")
	}
}
