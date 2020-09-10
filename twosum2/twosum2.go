package twosum2

import (
	. "lcode/entity"
)

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    tmpl := new(ListNode)
    l3 := tmpl
    carry := 0
    
    for (l1 != nil || l2 != nil || carry > 0){
        l3.Next = new(ListNode)
        l3 = l3.Next
        if (l1 != nil){
            carry += l1.Val
            l1 = l1.Next
        }
        if (l2 != nil){
            carry += l2.Val
            l2 = l2.Next
        }
        l3.Val = carry %10 
        carry = carry/10
	}
	
    return tmpl.Next
}
