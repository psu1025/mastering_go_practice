package cmdcalc

import (
	"fmt"
)

// EchoParameters function echo parameter when parameter until STOP.
func EchoParameters(vars []string) {
	for _, strvar := range vars {
		if strvar == "STOP" {
			break
		}
		fmt.Println(strvar)
	}
}
