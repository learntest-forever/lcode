package main

import (
	"fmt"
<<<<<<< HEAD
	"lcode/reverse2"
=======
	"lcode/reverse"
>>>>>>> 8d29411b7fbbbbb77d6f8c5a6a2685642f3e29d6
	"lcode/twosum"
)

func main() {
	l1 := []int{1, 3, 2}
	l2 := []int{2, 4, 6}
	twosum.Twosum(&l1, &l2)
<<<<<<< HEAD
	l3 := []int{1, 2, 5, 7}
	l5 := reverse2.Myreverse(l3, []int{})
	fmt.Printf("the l5 is :%v\n", l5[:])
=======
	l3 := []int{1, 2, 5}
	l4 := reverse.Myreverse(l3)
	fmt.Printf("the l4 is :%v\n", l4)
>>>>>>> 8d29411b7fbbbbb77d6f8c5a6a2685642f3e29d6
}
