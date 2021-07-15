/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {
	// 第一步要明确两点，「状态」和「选择」
	// 第二步要明确 dp 数组的定义。
	// 第三步，根据「选择」，思考状态转移的逻辑
	// return canPartition1(nums) //原始dp

	return canPartition2(nums) //进行状态压缩
}

// dp[i][j] = x 表示，对于前 i 个物品，当前背包的容量为 j 时，若 x 为 true，则说明可以恰好将背包装满，若 x 为 false，则说明不能恰好将背包装满。
// 根据这个定义，我们想求的最终答案就是 dp[N][sum/2]，base case 就是 dp[..][0] = true 和 dp[0][..] = false，因为背包没有空间的时候，就相当于装满了，
// 而当没有物品可选择的时候，肯定没办法装满背包。
func canPartition1(nums []int) bool {
	sum := 0

	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 { //和为奇数，不可能划分为两个相等集合的数组
		return false
	}

	n := len(nums)
	sum = sum / 2
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
	}

	//bad case
	for i := 0; i <= n; i++ {
		dp[i][0] = true //背包容量为0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] < 0 {
				//背包容量不足，不能装入第i个物品
				dp[i][j] = dp[i-1][j]
			} else {
				//装入或者不装入背包
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][sum]
}

func canPartition2(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}
	// C = half sum
	n, C, dp := len(nums), sum/2, make([]bool, sum/2+1)
	for i := 0; i <= C; i++ {
		dp[i] = (nums[0] == i)
	}
	for i := 1; i < n; i++ {
		for j := C; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[C]
}

// @lc code=end

