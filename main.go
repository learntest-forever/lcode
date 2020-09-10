package main

import (
	"fmt"
	"lcode/reverse2"
	"lcode/twosum"
)

func main() {
	l1 := []int{1, 3, 2}
	l2 := []int{2, 4, 6}
	twosum.Twosum(&l1, &l2)
	l3 := []int{1, 2, 5, 7}
	l5 := reverse2.Myreverse(l3, []int{})
	fmt.Printf("the l5 is :%v\n", l5[:])
}
