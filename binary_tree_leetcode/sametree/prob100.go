package sametreeleetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Base case: if both nodes are nil, they are the same
	if p == nil && q == nil {
		return true
	}

	// If one of the nodes is nil or the values are different, they are not the same
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	// Recursively check the left and right subtrees
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}
