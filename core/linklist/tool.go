package main


// convert *ListNode to []int
func l2s(head *Node) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// convert []int to *ListNode
func s2l(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	res := &Node{
		Val: nums[0],
	}

	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &Node{
			Val: nums[i],
		}

		temp = temp.Next
	}

	return res
}
