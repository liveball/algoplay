package my_sort

//归并排序思想：
//merge_sort(p…r) 表示，给下标从 p 到 r 之间的数组排序。我们将这个排序问题转化为了两个子问题，
//merge_sort(p…q) 和 merge_sort(q+1…r)，
//其中下标 q 等于 p 和 r 的中间位置，也就是 (p+r)/2。
//当下标从 p 到 q 和从 q+1 到 r 这两个子数组都排好序之后，
//我们再将两个有序的子数组合并在一起，这样下标从 p 到 r 之间的数据就也排好序了。
//归并排序算法，a是数组，n表示数组大小

// 时间复杂度：
//归并排序算法是一种在任何情况下时间复杂度都比较稳定的排序算法，
//这也使它存在致命的缺点，即归并排序不是原地排序算法，空间复杂度比较高，是 O(n)。
//正因为此，它也没有快排应用广泛。

func MergeSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}

	mergeSort(a, 0, n-1)
}

//递归函数
func mergeSort(a []int, p, r int) {
	//递归终止条件
	if p >= r {
		return
	}

	//取p到r之间的中间位置q
	q := (p + r) / 2
	//分治递归
	mergeSort(a, p, q)
	mergeSort(a, q+1, r)

	//讲a[p:q]和a[q+1:r] 合并为a[p:r]

	merge(a, p, q, r)
}

func merge(a []int, p, q, r int) {
	i := p
	j := q + 1
	k := 0 //临时数组索引

	//申请一个临时数组
	tArr := make([]int, r-p+1)
	for i <= q && j <= r {
		if a[i] <= a[j] {
			tArr[k] = a[i]
			i++
		} else {
			tArr[k] = a[j]
			j++
		}

		k++
	}

	//判断哪个子数组中有剩余数据
	start := i
	end := q
	if j <= r {
		start = j
		end = r
	}

	//将剩余数据拷贝到临时数组
	for start <= end {
		tArr[k] = a[start]
		k++
		start++
	}

	//将tArr 数组拷贝回到a[p:r]
	copy(a[p:r+1], tArr)
}
