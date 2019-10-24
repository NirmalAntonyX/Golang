package main

import (
	"fmt"
	"net/http"
)

func listIterator() ([]int16, bool) {
	var sliceVariable = make([]int16, 5)
	sliceVariable[0] = 34
	sliceVariable[1] = 35
	sliceVariable[2] = 36
	sliceVariable[3] = 37
	fmt.Println(sliceVariable)
	return sliceVariable, true
}
func main() {

	fmt.Println("Hello world")
	v, b := listIterator()

	if b {
		fmt.Println(v)
	}
	http.HandleFunc()

}
