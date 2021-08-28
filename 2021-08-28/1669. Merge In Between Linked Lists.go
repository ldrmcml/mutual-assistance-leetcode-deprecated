/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// algorithm:
// Critically, find preA node and postB node:
// 1. trick, dummy/fake head to simplify list head handling;1 <= a, maybe do not use dummy;
// 2. record pre-current node;
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{
		Next: list1,
	}

	// 1 <= a, maybe do not use dummy
	index := 0
	pre, preA, postB := list1, list1, list1
	for n := list1; n != nil; n = n.Next {
		if index == a {
			preA = pre
		}
		if index == b+1 {
			postB = n
		}
		pre = n
		index++
	}

	preA.Next = list2

	// obtain last node of list2
	for n := list2; n != nil; n = n.Next {
		pre = n
	}
	list2Last := pre

	list2Last.Next = postB

	return dummyHead.Next
}