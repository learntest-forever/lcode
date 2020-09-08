package twosum

import (
	"fmt"
	"strconv"
	"strings"
)

func Twosum(l1,l2 *[]int) *[]int{
	fmt.Println(*l1)
	fmt.Println(*l2)
	//翻转l1 l2
	l3 := reverse(*l1)
	fmt.Printf("the l3 is :%v\n",l3)	
	l4 := reverse(*l2)
	//slice to string,string to int
	l1ReversInt,_ := strconv.Atoi(slicetostr(l3))
	fmt.Printf("the l1ReversInt is :%v\n",l1ReversInt)
	l2ReversInt,_ := strconv.Atoi(slicetostr(l4))
	fmt.Printf("the l2ReversInt is :%v\n",l2ReversInt)
	l3New := l1ReversInt + l2ReversInt
	fmt.Printf("the l3New is :%v\n",l3New)
	l3StrOld := strconv.Itoa(l3New)
	l3Slice := strtoslice(l3StrOld)
	fmt.Printf("the l3Slice is :%v\n",l3Slice)
	target := reverse(l3Slice)
	//l3Reverse := reverse(l3Slice)
	//return &l3Reverse
	return &target
}

func reverse(s1 []int) []int{
	for from, to := 0, len(s1)-1; from < to; from, to = from+1, to-1 {
        s1[from], s1[to] = s1[to], s1[from]
    }
    return s1
}

//slice to string failed, todo fix
func slicetostr(s1 []int) string{
	s2 := strings.Replace(strings.Trim(fmt.Sprint(s1), "[]"), " ", "", -1)
	fmt.Printf("the s2 is :%s\n",s2)
	return s2
}

func strtoslice(s1 string) []int{
	s2 := []int{}
	for _,v := range s1 {
		//s2 = append(s2,strconv.Atoi(v))
		//fmt.Printf("strtoslice k is :%v, v is :%s\n",k,string(v))
		i3,_ := strconv.Atoi(string(v))
		s2 = append(s2,i3)
	}
	return s2
}