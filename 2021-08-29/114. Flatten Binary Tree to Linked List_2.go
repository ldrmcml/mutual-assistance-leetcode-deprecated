/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// in-place
// 1. 从根开始，左子树接到右孩子；
// 2. 将原来右子树接到最右边的右孩子；
// 3. 考虑新的右子树的根节点，重复，直到新的右子树为nil；
func flatten(root *TreeNode) {
	// 无论如何先判空！！！
	if root == nil {
		return
	}

	for root != nil {
		// 左子树为nil，直接考虑右节点，重复上述过程
		if root.Left == nil {
			root = root.Right
		} else {
			// 找左子树最右节点
			n := root.Left
			for ; n.Right != nil; n = n.Right {
			}
			n.Right = root.Right
			root.Right = root.Left
			// 切记将左子树置空
			root.Left = nil
			// 考虑新的右子树的根节点，重复上述过程
			root = root.Right
		}
	}
}