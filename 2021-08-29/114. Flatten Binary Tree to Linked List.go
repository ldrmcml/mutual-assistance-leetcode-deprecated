/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 暴力解法
// recursive? no way
// 1. convert to pre-traverse,
// 2. list to tree
var preArr []*TreeNode

func flatten(root *TreeNode) {
	preArr = nil
	preTraverse(root)
	convertToNewTree(root)
}

// recursive
func preTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	preArr = append(preArr, root)
	preTraverse(root.Left)
	preTraverse(root.Right)
}

func convertToNewTree(head *TreeNode) {
	pre := head
	for i, n := range preArr {
		if i == 0 {
			continue
		}
		pre.Left = nil
		pre.Right = n
		pre = n
	}
}