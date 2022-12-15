package lib

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepthWith(root *TreeNode) (i int) {
	if root == nil {
		return 0
	}
	list := []*TreeNode{root}
	for len(list) > 0 {
		i++
		nlist := []*TreeNode{}
		for _, v := range list {
			if v.Left != nil {
				nlist = append(nlist, v.Left)
			}
			if v.Right != nil {
				nlist = append(nlist, v.Right)
			}
		}
		list = nlist
	}
	return
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left) + 1
	rightMax := maxDepth(root.Right) + 1
	if leftMax < rightMax {
		leftMax = rightMax
	}
	return leftMax
}
