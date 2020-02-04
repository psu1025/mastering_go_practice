package cmdcalc

import "testing"

func Test_AvgAllVariable(t *testing.T) {
	numstrings := []string{"1", "2"}
	avg := AvgAllVariable(numstrings)

	if avg != 1.5 {
		t.Fatal("(1+2)/2 == 1.5")
	}
}

func Test_AvgAllVariable_Zero(t *testing.T) {
	numstrings := []string{}
	avg := AvgAllVariable(numstrings)

	if avg != 0.0 {
		t.Fatal("it must be zero")
	}
}
