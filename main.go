package main

import (
	"fmt"
	"lcode/reverse"
	"lcode/twosum"
)

func main() {
	l1 := []int{1, 3, 2}
	l2 := []int{2, 4, 6}
	twosum.Twosum(&l1, &l2)
	l3 := []int{1, 2, 5}
	l4 := reverse.Myreverse(l3)
	fmt.Printf("the l4 is :%v\n", l4)
}
