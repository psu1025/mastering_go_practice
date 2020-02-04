package cmdcalc

import (
	"fmt"
	"strconv"
)

// AddAllVariable is add all variable
func AddAllVariable(vars []string) (total int) {
	total = 0
	for _, num := range vars {
		i, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			continue
		}
		total += i
	}
	return
}
