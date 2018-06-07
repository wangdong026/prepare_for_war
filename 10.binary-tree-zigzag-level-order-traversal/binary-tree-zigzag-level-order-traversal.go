package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := initTree()
	ret := zigzagLevelOrder(node)

	for _, arr := range ret {
		for _, value := range arr {
			fmt.Print(value, "\n")
		}
		fmt.Println("===========")
	}
}

func initTree() *TreeNode {
	// 3,9,20,null,null,15,7
	ret := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}

	return ret
}

func zigzagLevelOrder(root *TreeNode) [][]int {

	res := [][]int{}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= len(res) {
			res = append(res, []int{})
		}

		if level%2 == 0 {
			// 正向
			res[level] = append(res[level], root.Val)
		} else {
			// 反向
			temp := make([]int, len(res[level])+1)
			temp[0] = root.Val
			copy(temp[1:], res[level])
			res[level] = temp
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)
	return res
}
