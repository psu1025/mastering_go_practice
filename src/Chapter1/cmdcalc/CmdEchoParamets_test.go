package cmdcalc

import "testing"

func Test_echoParameters(t *testing.T) {
	params := []string{"a", "b", "c", "STOP", "d"}
	EchoParameters(params)
}
