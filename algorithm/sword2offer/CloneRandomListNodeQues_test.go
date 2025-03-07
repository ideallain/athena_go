package sword2offer

import (
	common "athena_go/algorithm/common"
	"testing"
)

func CloneRandomListNode(head *common.RandomListNode) *common.RandomListNode {
	if head == nil {
		return nil
	}
	p1, p2 := head, head
	nodeMap := map[*common.RandomListNode]*common.RandomListNode{}
	for p1 != nil {
		nodeMap[p1] = &common.RandomListNode{Label: p1.Label}
		p1 = p1.Next
	}
	for p2 != nil {
		nodeMap[p2].Next = nodeMap[p2.Next]
		nodeMap[p2].Random = nodeMap[p2.Random]
		p2 = p2.Next
	}
	return nodeMap[head]
}

func TestCloneRandomListNode(t *testing.T) {
	randomListNode := common.RandomListNode{}
	randomListNode.Label = 1
	println("begin copy")
	resultNode := CloneRandomListNode(&randomListNode)
	print(resultNode)
}
