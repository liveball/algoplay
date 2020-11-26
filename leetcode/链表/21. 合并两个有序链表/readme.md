思路

我们可以如下递归地定义两个链表里的 merge 操作（忽略边界情况，比如空链表等）：

if list1[0]<list2[0]{
   list1[0]+merge(list1[1:],list2)
   list2[0]+merge(list1,list2[1:])
} else {
    list1[0]<list2[0]
}

也就是说，两个链表头部值较小的一个节点与剩下元素的 merge 操作结果合并。