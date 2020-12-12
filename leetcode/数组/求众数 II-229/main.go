package main

// 第一种比较简单，就是放入哈希表，最后将哈希表中计数大于n/3的键进行输出
func majorityElement(nums []int) []int {
    dic := make(map[int]int,0)
    threshold := len(nums) / 3
    var res []int
    for _,v := range nums{
        dic[v]++
    }
    for i,v := range dic{
        if v > threshold{
            res = append(res, i)
        }
    }
    return res
}

// 第二种，借助一个额外的哈希表，记录是否已经进行过输出，这样可以少遍历一次。但空间上有额外开销
func majorityElement2(nums []int) []int {
    dic := make(map[int]int,0)
    threshold := len(nums) / 3
    majDic := make(map[int]bool,0)
    var res []int
    for _,v := range nums{
        dic[v]++
        if dic[v] > threshold && majDic[v] == false{
            majDic[v] = true
            res = append(res,v)
        }
    }    
    return res
}


// 给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1)的算法解决此问题。
// 示例 1：
// 输入：[3,2,3]
// 输出：[3]
// 示例 2：

// 输入：nums = [1]
// 输出：[1]
// 示例 3：

// 输入：[1,1,1,3,3,2,2,2]
// 输出：[1,2]
// 提示：

// 1 <= nums.length <= 5 * 104
// -109 <= nums[i] <= 109
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/majority-element-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。