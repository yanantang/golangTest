package twentyDays

/* 876. 链表的中间结点
* https://leetcode-cn.com/problems/middle-of-the-linked-list/submissions/
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	one := head
	two := head
	for two != nil {
		if two.Next != nil {
			one = one.Next
		} else {
			break
		}
		two = two.Next
		if two.Next != nil {
			two = two.Next
		}
	}
	return one
}

func middleNodeV2(head *ListNode) *ListNode {
	one := head
	two := head
	for two != nil && two.Next != nil {
		one = one.Next
		two = two.Next.Next
	}
	return one
}

/* 19. 删除链表的倒数第 N 个结点
https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 1 先挪n步
	one, two := head, head
	for ; two != nil && n > 0; n-- {
		two = two.Next
	}

	if two == nil {
		return head.Next
	}

	// 2 同时走n步, 第一个的下一个就是要删除的节点
	for two.Next != nil {
		one = one.Next
		two = two.Next
	}

	// 3 删除节点
	del := one.Next
	if one.Next.Next == nil {
		one.Next = nil
	} else {
		one.Next = one.Next.Next
		del.Next = nil
	}

	return head
}

func removeNthFromEndV2(head *ListNode, n int) *ListNode {
	// 如果删除的是头节点, 用空节点占位
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	// 1 先挪n步
	one, two := dummy, head
	for i := 0; i < n; i++ {
		two = two.Next
	}

	// 2 同时走n步, 第一个的下一个就是要删除的节点
	for ; two != nil; two = two.Next {
		one = one.Next
	}
	one.Next = one.Next.Next
	return dummy.Next
}
