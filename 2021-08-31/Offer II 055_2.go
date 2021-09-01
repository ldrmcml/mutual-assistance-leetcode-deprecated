/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 1. 中序遍历将BST整颗树扁平化到数组；
// space: O(n), 均摊time: O(1)
type BSTIterator struct {
	arr []int
}

func (this *BSTIterator) inorder(root *TreeNode) {
	if root == nil {
		return
	}

	this.inorder(root.Left)
	this.arr = append(this.arr, root.Val)
	this.inorder(root.Right)
}

func Constructor(root *TreeNode) BSTIterator {
	b := BSTIterator{}
	b.inorder(root)
	return b
}

func (this *BSTIterator) Next() int {
	v := this.arr[0]
	this.arr = this.arr[1:]
	return v
}

func (this *BSTIterator) HasNext() bool {
	if len(this.arr) > 0 {
		return true
	}
	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */