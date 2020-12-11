package main

import(
	"fmt"
)

func main()  {
	deferCall()
}

// 打印中
// 打印前
// panic: 触发异常
// goroutine 1 [running]:

func deferCall()  {
	defer func ()  {
		fmt.Println("打印前")
	}()
	defer func ()  {
		fmt.Println("打印中")
	}()
	panic("触发异常")
	defer func ()  {
		fmt.Println("打印后")
	}()
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

// 给定一个头结点为 root 的链表, 编写一个函数以将链表分隔为 k 个连续的部分。

// 每部分的长度应该尽可能的相等: 任意两部分的长度差距不能超过 1，也就是说可能有些部分为 null。

// 这k个部分应该按照在链表中出现的顺序进行输出，并且排在前面的部分的长度应该大于或等于后面的长度。

// 返回一个符合上述规则的链表的列表。

// 举例： 1->2->3->4, k = 5 // 5 结果 [ [1], [2], [3], [4], null ]

// 示例 1：

// 输入: 
// root = [1, 2, 3], k = 5
// 输出: [[1],[2],[3],[],[]]
// 解释:
// 输入输出各部分都应该是链表，而不是数组。
// 例如, 输入的结点 root 的 val= 1, root.next.val = 2, \root.next.next.val = 3, 且 root.next.next.next = null。
// 第一个输出 output[0] 是 output[0].val = 1, output[0].next = null。
// 最后一个元素 output[4] 为 null, 它代表了最后一个部分为空链表。
// 示例 2：

// 输入: 
// root = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], k = 3
// 输出: [[1, 2, 3, 4], [5, 6, 7], [8, 9, 10]]
// 解释:
// 输入被分成了几个连续的部分，并且每部分的长度相差不超过1.前面部分的长度大于等于后面部分的长度。
//  

// 提示:

// root 的长度范围： [0, 1000].
// 输入的每个节点的大小范围：[0, 999].
// k 的取值范围： [1, 50].


// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/split-linked-list-in-parts
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。