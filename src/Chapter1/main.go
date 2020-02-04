package main

import (
	"fmt"
	"os"

	"./cmdcalc"
)

func main() {
	argsWithoutProg := os.Args[1:]
	fmt.Println("add all cmd variables")
	fmt.Println(cmdcalc.AddAllVariable(argsWithoutProg))

	fmt.Println("avg all cmd variables")
	fmt.Println(cmdcalc.AvgAllVariable(argsWithoutProg))
}
