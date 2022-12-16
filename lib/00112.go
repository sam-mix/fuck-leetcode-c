package lib

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Val == targetSum && root.Right == nil && root.Left == nil {
		return true
	}
	newTargetSum := targetSum - root.Val
	return hasPathSum(root.Left, newTargetSum) || hasPathSum(root.Right, newTargetSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Right == nil && root.Left == nil {
		return root.Val == targetSum
	}
	newTargetSum := targetSum - root.Val
	return hasPathSum(root.Left, newTargetSum) || hasPathSum(root.Right, newTargetSum)
}
