package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := makeListNode([]int{1, 8})
	l2 := makeListNode([]int{0})

	ret := addTwoNumbers(l1, l2)

	for {
		println(ret.Val)

		ret = ret.Next
		if ret == nil {
			break
		}
	}
}

func makeListNode(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}

	res := &ListNode{
		Val: is[0],
	}
	temp := res

	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}

	return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	ret := &ListNode{}
	tmp := ret
	v, n := 0, 0 // v :当前位相加结果 n : 进位值

	for {
		//println("l1:", l1.Val, " l2:", l2.Val)
		v, n = add(l1, l2, n)

		//println("v:", v, " n:", n)
		tmp.Val = v

		l1 = next(l1)
		l2 = next(l2)

		if l1 == nil && l2 == nil {
			break
		}

		tmp.Next = &ListNode{}
		tmp = tmp.Next
	}
	if n == 1 {
		tmp.Next = &ListNode{Val: 1, Next: nil}
	}
	return ret
}

func next(l *ListNode) (*ListNode) {
	if l != nil {
		return l.Next
	}
	return nil
}

func add(i *ListNode, j *ListNode, lastN int) (v, n int) {
	sum := 0
	if i != nil {
		sum += i.Val
	}
	if j != nil {
		sum += j.Val
	}
	sum += lastN

	n = 0
	v = sum

	if sum >= 10 {
		v = sum - 10
		n = 1
	}
	return v, n
}
