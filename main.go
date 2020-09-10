package main

import (
	"fmt"
	"lcode/findsum"
)

func main() {
	s1 := []int{1, 5, 8, 11}
	target := 13
	s2 := findsum.Findsum(s1, target)
	fmt.Println(s2)
}
