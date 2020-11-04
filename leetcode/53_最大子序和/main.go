package main

import "fmt"

var (
	numsTest = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
)

func main() {
	sum := sumMax(numsTest)
	sum2 := maxSumDP(numsTest)
	fmt.Println(sum, sum2)
}

//给定一个数组nums， 找到一个具有最大的连续子数组(子数组最少包含一个元素)。返回其最大和

//题目解析：要求子序最大和，得循环累加求解
func sumMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0] //假设第一个元素最大
	sum := 0       //计算和
	for i := 1; i < len(nums)-1; i++ {
		if sum > 0 { //大于0 则累加
			sum = sum + nums[i]
		} else { //小于0 则覆盖
			sum = nums[i]
		}
		res = max(res, sum) //每次循环之后获取最大的和
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//构造状态转移方程：
// if dp[i-1]>=0 dp[i]= dp[i-1] +nums[i]
// else dp[i-1]<0 dp[i]  nums[i]
func maxSumDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0] //存第一个元素为最大序列和

	res := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if dp[i-1] >= 0 { //上个序列和大于0则累加
			dp[i] = dp[i-1] + nums[i]
		} else { //否则覆盖
			dp[i] = nums[i]
		}

		res = max(res, dp[i])
	}
	return res
}
