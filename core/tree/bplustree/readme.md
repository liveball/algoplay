
###参考文章
>1.漫画：什么是B-树？https://zhuanlan.zhihu.com/p/54084335  
>2.漫画：什么是B+树？https://zhuanlan.zhihu.com/p/54102723   

###B+树特点
>1.每一个父节点的元素都出现在子节点中，是子节点的最大（或者最小）的元素  
>2.每个节点中子节点的个数不能超过 m，也不能小于 m/2；  
>3.根节点的子节点个数可以不超过 m/2，这是一个例外；  
>4.m 叉树只存储索引，并不真正存储数据，这个有点儿类似跳表；  
>5.通过链表将叶子节点串联在一起，这样可以方便按区间查找；  
>6.一般情况，根节点会被存储在内存中，其他节点存储在磁盘中  

```p 除了 B+ 树，你可能还听说过 B 树、B- 树，我这里简单提一下。
实际上，B- 树就是 B 树，英文翻译都是 B-Tree，
这里的“-”并不是相对 B+ 树中的“+”，而只是一个连接符。
这个很容易误解，所以我强调下。
而 B 树实际上是低级版的 B+ 树，或者说 B+ 树是 B 树的改进版。
```

### B 树跟 B+ 树的不同点主要集中在这几个地方：
>1.B+ 树中的非叶子节点不存储数据，只是索引   
>2.B 树中的节点存储数据  
>3.B 树中的叶子节点并不需要链表来串联    

### B+树的优势：
>1.单一节点存储更多的元素，使得查询的IO次数更少  
>2.所有查询都要查找到叶子节点，查询性能稳定  
>3.所有叶子节点形成有序链表，便于范围查询  