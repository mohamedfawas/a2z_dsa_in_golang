package binarytreeinordertraversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	// Result slice to store the traversal
	var result []int

	// Helper function to perform inorder traversal recursively
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		// Traverse the left subtree
		inorder(node.Left)
		// Visit the current node
		result = append(result, node.Val)
		// Traverse the right subtree
		inorder(node.Right)
	}

	// Start traversal from the root
	inorder(root)
	return result
}
