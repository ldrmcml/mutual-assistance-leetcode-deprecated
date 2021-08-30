/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 1. 找到中点位置；
// 1. 原地逆转后半段；
// 2. 原链表与逆转后链表归并：间隔插入，直到遇到相同节点；
func reorderList(head *ListNode) {
	// 0,1,2 nodes, return directly
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	mid := findMid(head)
	h2 := reverse(mid)
	merge(head, h2)
}

// 原来链表最后指向nil，后半段最后同样指向nil
func findMid(head *ListNode) *ListNode {
	s, f := head, head
	for ; f != nil && f.Next != nil; s, f = s.Next, f.Next.Next {

	}

	mid := s.Next
	s.Next = nil

	return mid
}
func reverse(mid *ListNode) *ListNode {
	if mid == nil || mid.Next == nil {
		return mid
	}

	var pre *ListNode
	for node := mid; node != nil; {
		tmp := node.Next
		node.Next = pre
		pre = node
		node = tmp
	}

	return pre
}
func merge(h1, h2 *ListNode) {
	n := new(ListNode)
	for i := 0; h1 != nil && h2 != nil; i++ {
		if i%2 == 0 {
			n.Next = h1
			h1 = h1.Next
		} else {
			n.Next = h2
			h2 = h2.Next
		}
		// holy shit!!!
		n = n.Next
	}

	if h1 != nil {
		n.Next = h1
	}
	if h2 != nil {
		n.Next = h2
	}
}