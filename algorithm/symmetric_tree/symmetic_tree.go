package main

// https://leetcode.com/problems/symmetric-tree/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricRec(root.Left, root.Right)
}

// 단순 값만 비교
func isSymmetricRec(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSymmetricRec(p.Left, q.Right) && isSymmetricRec(p.Right, q.Left)
}

func main() {
	tree := &TreeNode{Val: 0}
	tree.Left = &TreeNode{Val: 1}
	tree.Left.Left = &TreeNode{Val: 2}
	tree.Right = &TreeNode{Val: 1}
	tree.Right.Right = &TreeNode{Val: 2}

	if isSymmetric(tree) == false {
		panic("This must be true.")
	}

}
