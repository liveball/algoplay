package sort

func QuickSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	quickSort(nums, 0, n-1)
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	m := partition(nums, l, r)

	quickSort(nums, l, m-1)
	quickSort(nums, m+1, r)

}

func partition(nums []int, l, r int) int {
	pivot := nums[r] //假设中心点数据为当前数组第r个元素
	nl := l

	for i := l; i < r; i++ {
		if nums[i] < pivot { //如果当前元素小于中心点数据，则将 中心点左边元素和当前元素进行交换, nl+1
			nums[i], nums[l] = nums[l], nums[i]
			nl++
		}
	}

	nums[nl], nums[r] = nums[r], nums[nl]

	return nl

}
