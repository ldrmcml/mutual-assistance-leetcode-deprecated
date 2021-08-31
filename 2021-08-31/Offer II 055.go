/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 1. 使用栈保存节点指针，初始化时仅从根节点一左到底，没有完全中序遍历整颗二叉树；
// 2. Next方法：弹出栈顶，判断是否存在右子树，如果存在，再次一左到底；然后返回栈顶对应val；
type BSTIterator struct {
	arr []*TreeNode
}

func Constructor(root *TreeNode) (b BSTIterator) {
	if root == nil {
		return b
	}
	// 一路左子树到底
	for ; root != nil; root = root.Left {
		b.arr = append(b.arr, root)
	}
	return b
}

func (this *BSTIterator) Next() int {
	// 读取栈顶
	n := len(this.arr)
	cur := this.arr[n-1]
	// 弹出栈顶，然后push右节点，中序：左中右
	this.arr = this.arr[:n-1]
	if cur.Right != nil {
		// 一路左子树到底
		for node := cur.Right; node != nil; node = node.Left {
			this.arr = append(this.arr, node)
		}
	}
	return cur.Val
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