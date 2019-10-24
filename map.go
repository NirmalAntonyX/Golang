package main

import (
	"fmt"
	"sort"
)

type mapi map[string]int

func main() {
	myMap := mapi{"nat": 6,
		"max":  1,
		"shy":  38,
		"nirm": 38}

	myMap["san"] = 40
	myMap["leen"] = 34

	delete(myMap, "nirm")
	fmt.Println(myMap)

	intlist := make([]int, 100)
	i := 0
	for _, value := range myMap {
		intlist[i] = value
		i += 1
	}
	fmt.Println(intlist, sort.IntsAreSorted(intlist))

	sort.IntSlice(intlist).Sort()

	fmt.Println(intlist, sort.IntsAreSorted(intlist))
}
