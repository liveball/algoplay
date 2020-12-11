package ten_sort

func insertSort(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	for i, v := range nums {
		preIndex := i - 1
		cur := v

		for preIndex >= 0 && nums[preIndex] > cur { //比较当前值和已排序区间逐一进行比较，寻找合适的位置
			nums[preIndex+1] = nums[preIndex] //腾出插入位置
			preIndex--
		}

		nums[preIndex+1] = cur
	}
}
