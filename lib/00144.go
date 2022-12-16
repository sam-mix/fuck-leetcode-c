package lib

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) (ans []int) {
	var pre func(*TreeNode)
	pre = func(root *TreeNode) {
		if root != nil {
			ans = append(ans, root.Val)
			pre(root.Left)
			pre(root.Right)
		}
	}
	pre(root)
	return ans
}
