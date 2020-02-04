package main

import "fmt"

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	var idMap = make(map[int]string)
	idMap[0] = "wow"
	idMap[3] = "wow"

	fmt.Println(idMap)
	delete(idMap, 3)
	fmt.Println(idMap)

	val, exists := idMap[4]
	fmt.Println(val, exists)

	val2, exists2 := idMap[0]
	fmt.Println(val2, exists2)
}
