package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node := initTree()
	println(maxDepth(node))
}

func initTree() *TreeNode {
	// 3,9,20,null,null,15,7
	ret := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}

	return ret
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
