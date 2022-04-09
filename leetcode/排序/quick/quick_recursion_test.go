package quick

import (
	"testing"
)

var (
	nums = []int{
		3, 2, 5, 1, 6,
	}
)

func Test_qsort(t *testing.T) {
	qsort(nums, 0, len(nums)-1)
	t.Log(nums)
}

func Test_qsort2(t *testing.T) {
	qsort2(nums, 0, len(nums)-1)
	t.Log(nums)
}

//递归方式
func qsort(arr []int, left, right int) {
	if len(arr) == 0 || left < 0 || right <= 0 || left >= right {
		return
	}

	k := partition(arr, left, right)
	if k > left {
		qsort(arr, left, k-1)
	}

	if k < right {
		qsort(arr, k+1, right)
	}
}

//非递归方式
func qsort2(arr []int, left, right int) {
	if len(arr) == 0 || left < 0 || right <= 0 || left >= right {
		return
	}

	stack := make([]int, 0, len(arr))
	//（注意保存顺序）先将初始状态的左右指针压栈
	stack = append(stack, right) //先存右指针
	stack = append(stack, left)  //再存左指针
	for len(stack) != 0 {
		i := stack[len(stack)-1] //先弹出左指针
		stack = stack[:len(stack)-1]
		j := stack[len(stack)-1] //再弹出右指针
		stack = stack[:len(stack)-1]

		if i < j {
			k := partition(arr, i, j)
			if k > i {
				stack = append(stack, k-1) //保存中间变量
				stack = append(stack, i)   //保存中间变量
			}
			if j > k {
				stack = append(stack, j)   //保存中间变量
				stack = append(stack, k+1) //保存中间变量
			}
		}
	}

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
