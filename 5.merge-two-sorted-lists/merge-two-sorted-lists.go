package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// 1->2->4, 1->3->4
	l1 := buildListNode([]int{1, 2, 4})
	l2 := buildListNode([]int{1, 3, 4})
	result := mergeTwoListNode(l1, l2)
	for {
		println(result.Val)
		if result.Next == nil {
			break
		}
		result = result.Next
	}
}

func buildListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	ret := &ListNode{}
	tmp := ret
	for _,num := range nums {
		tmp.Next = &ListNode{Val:num}
		tmp = tmp.Next
	}

	return ret.Next
}

func mergeTwoListNode(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	head := &ListNode{}
	pointSentry := head

	for {
		if l1.Val <= l2.Val {
			pointSentry.Next = l1
			l1 = l1.Next
		}else {
			pointSentry.Next = l2
			l2 = l2.Next
		}
		pointSentry = pointSentry.Next
		if l1 == nil || l2 == nil{
			break
		}
	}

	if l1 != nil {
		pointSentry.Next = l1
	}
	if l2 != nil{
		pointSentry.Next = l2
	}
	return head.Next
}
