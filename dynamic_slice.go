package main

import (
	"fmt"
)

func oddOccur(A []int) {

	var histogram = make(map[int]int)

	for _, value := range A {
		histogram[value]++
	}
	for key, value := range histogram {
		fmt.Println("key:%d value:%v size:%d", key, value, len(histogram))
	}
}
func main() {

	a := []int{3, 6, 9, 96, 3, 4}
	oddOccur(a)

	b := -1 << 31
	fmt.Println("0x%x", b)
}
