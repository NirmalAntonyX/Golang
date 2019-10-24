package main

import (
	"fmt"
	"sort"
)

func sortFunction(m map[int]int) (indexes []int) {

	for _, val := range m {
		indexes = append(indexes, val)
	}
	sort.Ints(indexes)
	return

}
func main() {

	mapElement := map[int]int{
		1: 3,
		2: 2,
		3: 1,
	}
	sortFunction(mapElement)

	for k, val := range mapElement {
		fmt.Println("key: ", k, "value : ", val)
	}

}
