package test

func lengthOfLIS(nums []int) int {

	// return lengthOfLIS_Recursion(nums)// 递归 时间复杂度 O(2^n)

	//return dp_BottomUp(nums) //自底向上 遍历子问题. 时间复杂度: O(n^2), 空间复杂度: O(n)

	return lengthOfLIS_BinarySearch(nums) //二分查找插入位置 时间复杂度: O(n*log(n)), 空间复杂度: O(n)
}

//自底向上指的是，通过状态转移方程，从最小的问题规模入手，不断的增加问题规模，
//直到处理完成所有的问题规模，依然使用记忆化避免重复计算，不需要递归
func dp_BottomUp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	res := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1 //初始化 所有以当前元素结尾的最长序列长度为1
	}

	//     递推式为 dp[i] = max(dp[i], dp[j]+1) for j from 0 to i,其含义为:
	//     以nums[i]为结束的最长上升子序列， 是之前所有子序列中最长上升子序列长度+1的最大值

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] { // nums[i]与i前面的数nums[j]一一对比，如果nums[i]比较大，则候选长度+1
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		res = max(res, dp[i]) // 更新当前最长递增子序列的长度
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS_BinarySearch(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lis := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ { //在长度为length的数组中， 找到nums[i]的插入位置
		pos := binarySearch(lis, 0, len(lis)-1, nums[i])

		if pos == len(lis) { //如果pos等于当前长度， 说明i已追加到数组d的尾部
			lis = append(lis, nums[i])
		} else {
			lis[pos] = nums[i] //覆盖
		}
	}

	return len(lis)
}

func binarySearch(nums []int, low, high, target int) int {
	for low <= high {
		mid := (low + high) >> 1

		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

var maxLen = 1

func lengthOfLIS_Recursion(nums []int) int {
	lis(nums, len(nums))
	return maxLen
}

// 时间复杂度
// 用公式法解决该递归问题的时间复杂度，如下。

// 当 n=1 的时候，递归直接返回 1，执行时间为 O(1)，即 T(1)=O(1)
// 当 n=2 的时候，内部调用了一次递归求解 T(1)，所以 T(2)=T(1)
// 当 n=3 的时候，T(3)=T(1) + T(2)
// …

// T(n−1)=T(1) + T(2) + … + T(n−2)
// T(n)=T(1) + T(2) + … + T(n−1)
// 通过观察，我们得到：T(n)=2×T(n−1)，这并不满足 T(n)=a×T(n / b) + f(n) 的关系式

// 但是 T(n) 等于两倍的 T(n−1)，表明，我们的计算是成指数增长的，每次的计算都是先前的两倍。
// 所以 O(n)=O(2n)

func lis(nums []int, n int) int {
	if n <= 1 {
		maxLen = n
		return n
	}

	res := 0
	maxEnding := 1
	for i := 1; i < n; i++ {
		res = lis(nums, i) //递归调用函数

		if nums[i-1] < nums[n-1] && res+1 > maxEnding {
			maxEnding = res + 1
		}
	}

	if maxLen < maxEnding { //如果那个数比目前最后一个数都要小， 那么就构成一个新的上升子序列
		maxLen = maxEnding
	}

	return maxEnding //返回以当前数字结尾的上升子序列的最大长度
}
