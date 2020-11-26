package main

func main()  {
	
}

/**
 * Definition for singly-linked list.
 */

 type ListNode struct {
      Val int
     Next *ListNode
 }
 
 func splitListToParts(root *ListNode, k int) []*ListNode {
	cur:=root
	l:=0
	for cur!=nil{//求链表长度
		cur=cur.Next
		l++
	}
	
	 avg:=l/k //每个子链表的平均元素的个数 (有多少个长度为 (avg + 1) 的子链表排在前面)
	 remain:=l%k //剩下的个节点不能丢，得全部塞到子链表里面去
	 res:=make([]*ListNode, k)

	 head:=root
	 var pre *ListNode
	 for i:=0; i<k; i++{//数组有k个元素需要遍历k次
	   res[i] =head //长的链表排前面，短的链表排后面，
	   
	   var tl int 
	   if remain>0{
		   tl=avg+1//只有前面的两个子链表分别分担一个多余的节点
	   } else {
		   tl=avg
	   }

	   for j:=0; j<tl; j++{
		   pre = head
		   head = head.Next
	   }
		
		if pre!=nil{//一个子链表已经生成，断开连接
			pre.Next = nil
		}

		if remain>=1{//只要还有剩余就表示还要塞子链表
			remain--
		} 
	 }

	 return res

}