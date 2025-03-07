package sword2offer

import common "athena_go/algorithm/common"

func FindKthToTail(head *common.ListNode, k int) *common.ListNode {
	// write code here
	if head == nil || k <= 0 {
		return nil
	}
	node1, node2 := head, head
	for i := 1; i < k; i++ {
		node2 = node2.Next
		if node2 == nil {
			return nil
		}
	}
	for node2.Next != nil {
		node1 = node1.Next
		node2 = node2.Next
	}
	return node1
}
