package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{
		Val:  0,
		Next: head,
	}

	pre := dummyNode
	cur := head
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			pre = pre.Next
		}
		i++
	}

	pre.Next = pre.Next.Next
	return dummyNode.Next
}
