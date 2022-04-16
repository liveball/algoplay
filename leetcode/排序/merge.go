package test

func mergeSort(a []int) {
	mergeSortA(a, 0, len(a)-1)
}

func mergeSortA(a []int, p int, r int) {
	if p >= r {
		return
	}
	//取p,r 中间的位置q
	q := (p + r) / 2
	//分治递归
	mergeSortA(a, p, q)
	mergeSortA(a, q+1, r)

	//将a[p...q] 和 a[q+1...r] 合并为a[p...r]

	merge(a, p, q, r)
}

func merge(a []int, p, q, r int) {
	i, j := p, q+1
	k := 0
	tmp := make([]int, r-p+1)

	for i <= q && j <= r {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
		k++
	}

	//查看a[p...q] 剩余
	for ; i <= q; i++ {
		tmp[k] = a[i]
		k++
	}
	//查看a[q+1...r] 剩余
	for ; j <= r; j++ {
		tmp[k] = a[j]
		k++
	}

	//拷贝 tmp[p:r-p] 到a[p:r]
	//for n := p; n <= r-p; n++ {
	//	a[n] = tmp[n]
	//}

	copy(a[p:r+1], tmp)
}
