package test

func quickSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	quickSortA(nums, 0, n-1)
}

func quickSortA(nums []int, l, r int) {
	if l >= r {
		return
	}

	p := partition(nums, l, r)

	quickSortA(nums, l, p-1)
	quickSortA(nums, p+1, r)

}

func partition(nums []int, l, r int) int {
	pivot := nums[r] //假设分区点数据为当前数组第r个元素
	nl := l

	for i := l; i < r; i++ {
		if nums[i] < pivot { //如果当前元素小于分区点数据，则将 分区点左边元素和当前元素进行交换, nl+1
			nums[i], nums[l] = nums[l], nums[i]
			nl++
		}
	}

	nums[nl], nums[r] = nums[r], nums[nl] //交换分区点

	return nl

}
