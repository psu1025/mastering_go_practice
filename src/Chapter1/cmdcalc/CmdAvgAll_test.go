package cmdcalc

import "testing"

func Test_AvgAllVariable(t *testing.T) {
	numstrings := []string{"1", "2"}
	sum := AvgAllVariable(numstrings)

	if sum != 1.5 {
		t.Fatal("(1+2)/2 == 1.5")
	}
}

func Test_AvgAllVariable_Zero(t *testing.T) {
	numstrings := []string{}
	sum := AvgAllVariable(numstrings)

	if sum != 0 {
		t.Fatal("it must be zero")
	}
}
