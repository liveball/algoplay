package ten_sort

func selectSort(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	//从无序区间不断挑选最小值，挑选n-1次（最后一个元素不用挑选）
	for i := 0; i < length-1; i++ {
		//i下标是已经排好序的元素，右边(包括i)是无序区间
		//最小值默认是无序区间的第一个元素下标(也就是i)
		min := i

		for j := i + 1; j < length; j++ { //遍历无序区间除min外的元素
			if nums[j] < nums[min] { //如果遇到比默认值小的元素
				min = j //更新最小值的下标
			}
		}

		//交换最小值与无序区间的第一个元素
		nums[i], nums[min] = nums[min], nums[i]
	}
}
