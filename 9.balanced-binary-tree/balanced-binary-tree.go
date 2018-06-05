package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	testTree := initTree()
	println(isBalanced(testTree))
}

func initTree() *TreeNode {
	// 3,9,20,null,null,15,7
	ret := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}

	return ret
}

func isBalanced(root *TreeNode) bool {
	_, isBalance := getNodeDepth(root)
	return isBalance
}

func getNodeDepth(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	leftDepth, leftIsBalance := getNodeDepth(node.Left)
	rightDepth, rightIsBalance := getNodeDepth(node.Right)

	if leftIsBalance && rightIsBalance && leftDepth-rightDepth <= 1 && leftDepth-rightDepth >= -1 {
		max := leftDepth
		if rightDepth > leftDepth {
			max = rightDepth
		}

		return max + 1, true
	}
	return 0, false
}
