package cmdcalc

import "testing"

func Test_AddAllVariable(t *testing.T) {
	numstrings := []string{"1", "2", "3", "ㄱ"}
	sum := AddAllVariable(numstrings)

	if sum != 6 {
		t.Fatal("1+2+3 == 6")
	}
}
