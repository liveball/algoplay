package main

import "fmt"

//快排的思想是这样的：如果要排序数组中下标从 p 到 r 之间的一组数据，
//我们选择 p 到 r 之间的任意一个数据作为 pivot（分区点）。
//我们遍历 p 到 r 之间的数据，将小于 pivot 的放到左边，
//将大于 pivot 的放到右边，将 pivot 放到中间。
//
//经过这一步骤之后，数组 p 到 r 之间的数据就被分成了三个部分，
//前面 p 到 q-1 之间都是小于 pivot 的，
//中间是 pivot，后面的 q+1 到 r 之间是大于 pivot 的

//时间复杂度：

//快速排序算法虽然最坏情况下的时间复杂度是 O(n2)，
//但是平均情况下时间复杂度都是 O(nlogn)。
//不仅如此，快速排序算法时间复杂度退化到 O(n2) 的概率非常小，
//我们可以通过合理地选择 pivot 来避免这种情况。

//递推公式：
//quick_sort(p…r) = quick_sort(p…q-1) + quick_sort(q+1… r)
//
//终止条件：
//p >= r

func QuickSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}

	partition2(a, 0, n-1)
}

func quickSort(a []int, p, r int) {
	if p >= r {
		return
	}

	q := partition(a, p, r)
	quickSort(a, p, q-1)
	quickSort(a, q+1, r)
}

func partition2(a []int, p, r int) int {
	// 选取最后一位当对比数字
	pivot := a[r]

	i := p

	fmt.Println("pivot", pivot, i, r)
	for j := p; j < r; j++ {
		fmt.Println("before exchange by pivot", a)
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
		fmt.Println("after exchange by pivot", a)
	}

	a[i], a[r] = a[r], a[i]

	fmt.Println("pivot", i, a)
	return i
}

//partition() 分区函数实际上我们前面已经讲过了，
//就是随机选择一个元素作为 pivot（一般情况下，
// 可以选择 p 到 r 区间的最后一个元素），
//然后对 A[p…r]分区，函数返回 pivot 的下标。
//如果我们不考虑空间消耗的话，partition() 分区函数可以写得非常简单。
//我们申请两个临时数组 X 和 Y，遍历 A[p…r]，
//将小于 pivot 的元素都拷贝到临时数组 X，
//将大于 pivot 的元素都拷贝到临时数组 Y，
//最后再将数组 X 和数组 Y 中数据顺序拷贝到 A[p…r]。

//但是，如果按照这种思路实现的话，
//partition() 函数就需要很多额外的内存空间，
//所以快排就不是原地排序算法了。
//如果我们希望快排是原地排序算法，
//那它的空间复杂度得是 O(1)，
//那 partition() 分区函数就不能占用太多额外的内存空间，
//我们就需要在 A[p…r]的原地完成分区操作。

//这里的处理有点类似选择排序。我们通过游标 i 把 A[p…r]分成两部分。
//A[p…i]的元素都是小于 pivot 的，
//我们暂且叫它“已处理区间”，A[i+1…r]是“未处理区间”。
//我们每次都从未处理的区间 A[i+1…r]中取一个元素 A[j]，
//与 pivot 对比，如果小于 pivot，则将其加入到已处理区间的尾部，
//也就是 A[i]的位置
