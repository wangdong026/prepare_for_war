package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := buildListNode([]int{1, 2})
	result := removeNthFromEnd(l, 1)
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
	for _, num := range nums {
		tmp.Next = &ListNode{Val: num}
		tmp = tmp.Next
	}

	return ret.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	p2 := head
	for p2 != nil {
		if n < 0 {
			p1 = p1.Next
		}
		p2 = p2.Next
		n--
	}

	if n >= 0 {
		return head.Next
	}
	p1.Next = p1.Next.Next

	return head
}
