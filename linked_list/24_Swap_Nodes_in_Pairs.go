package linked_list


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	temp := head.Next
	head.Next = head.Next.Next
	temp.Next = head
	head.Next = swapPairs(head.Next)
	return temp
}
