/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */

// @lc code=start
func peakIndexInMountainArray(arr []int) int {
	return peakIndexInMountainArray0(arr)

	// return peakIndexInMountainArray1(arr)

	// return peakIndexInMountainArray2(arr)
}

func peakIndexInMountainArray0(arr []int) int {
	for i := 1; ; i++ {
		if arr[i] > arr[i+1] {
			return i
		}
	}
}

func peakIndexInMountainArray1(A []int) int {
	low, high := 0, len(A)-1
	for low < high {
		mid := low + (high-low)>>1
		if A[mid] > A[mid+1] && A[mid] > A[mid-1] {
			return mid
		}

		if A[mid] > A[mid+1] && A[mid] < A[mid-1] {
			high = mid - 1
		}

		if A[mid] < A[mid+1] && A[mid] > A[mid-1] {
			low = mid + 1
		}
	}

	return 0
}

func peakIndexInMountainArray2(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1

		// 如果 mid 较大，则左侧存在峰值，high = m，
		// 如果 mid + 1 较大，则右侧存在峰值，low = mid + 1
		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

// @lc code=end

