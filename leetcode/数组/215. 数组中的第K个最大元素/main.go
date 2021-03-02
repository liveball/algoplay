package main

import "fmt"

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))          //5
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) //4

	fmt.Println(findKthLargestMaxHeap([]int{3, 2, 1, 5, 6, 4}, 2))          //5
	fmt.Println(findKthLargestMaxHeap([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) //4
}

// func findKthLargest(nums []int, k int) int {
//     sort.Ints(nums)
//     fmt.Println(nums)
//     return nums[len(nums)-k]
// }

func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(nums []int, l, r, k int) int {
	p := partition(nums, l, r)
	if p == k {
		return nums[p]
	} else if p < k { //在分区点右边
		return quickSort(nums, p+1, r, k)
	} else { //在分区点左边
		return quickSort(nums, l, p-1, k)
	}
}

func partition(nums []int, l, r int) int {
	pivot := nums[r] //假设最右边的元素的索引为分区点
	i := l           //记录最左边索引
	for j := l; j < r; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++ //移动
		}
	}

	nums[i], nums[r] = nums[r], nums[i]
	return i
}

//构造大顶堆，做k-1次删除，对顶元素就是要找的第k大元素
func findKthLargestMaxHeap(nums []int, k int) int {
	heapSize := len(nums)
	//构造大顶堆,使得堆顶最大元素是nums[0]
	buildMaxHeap(nums, heapSize)

	//删除k-1次堆顶元素，注意每次循环len(nums) 会减少一个元素
	for i := len(nums) - 1; i > len(nums)-k; i-- {
		//交换堆顶元素和最后一个节点的元素，使堆顶元素交换到最后，数组长度减1删除
		nums[0], nums[i] = nums[i], nums[0]
		//堆元素数量减1
		heapSize--
		//每次从堆顶元素开始堆化
		heapify(nums, 0, heapSize)
	}

	return nums[0] //返回堆顶元素
}

func buildMaxHeap(nums []int, heapSize int) {
	for i := len(nums) / 2; i >= 0; i-- {
		heapify(nums, i, heapSize)
	}
}

func heapify(nums []int, i, heapSize int) {
	l := 2*i + 1
	r := 2*i + 2
	max := i

	if l < heapSize && nums[max] < nums[l] {
		max = l
	}
	if r < heapSize && nums[max] < nums[r] {
		max = r
	}

	if i != max {
		//交换堆顶元素为最大的元素
		nums[i], nums[max] = nums[max], nums[i]
		//递归堆化
		heapify(nums, max, heapSize)
	}
}

// 215. 数组中的第K个最大元素
// 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

// 示例 1:

// 输入: [3,2,1,5,6,4] 和 k = 2
// 输出: 5
// 示例 2:

// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
// 输出: 4
// 说明:

// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
