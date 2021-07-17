/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	l := len(nums)
	res := make([][]int, 0)
	if l == 0 {
		return res
	}

	path := []int{} // path跟visited是全过程共享的，一定要注意恢复状态
	visited := make(map[int]bool)
	dfs(nums, l, 0, path, visited, &res)
	return res

	// return permuteRecursion(nums)//递归
}

// 遍历决策树中所有可行路径，并保存结果。
// 如果在递归调用传递参数时，传递同一个底层数组的引用的话，那一定要注意两点：

// 在保存结果时要保存当前状态（path）的快照，而不是path本身，因为状态是共享的，会被后续修改
// 修改状态->递归调用->调用返回，调用返回后要恢复状态，保证每一层的不同决策分支的初始状态不变
// 另一种做法是在每次递归调用时就传递一份当前状态的拷贝

func dfs(nums []int, l int, depth int, path []int, visited map[int]bool, res *[][]int) {
	if depth == l {
		// 保存path当前快照，因为path是共享的，后续会被其他决策修改
		*res = append(*res, append([]int(nil), path...)) //方法一

		// *res = append(*res, path) //方法二
		return
	}

	for i, num := range nums {
		if !visited[i] {
			visited[i] = true // 更改当前状态，递归调用下一层

			//方法一
			path = append(path, num)
			dfs(nums, l, depth+1, path, visited, res)
			path = path[:len(path)-1] //撤销当前已使用节点

			// 方法二中，因为每次递归调用传的都是新slice，path本身是不会被修改的，
			// 只是在初始path的基础上每一层调用添加新元素。
			// 每个函数调用栈内的局部变量path都是一个独立的path，
			// 是不会变的。append之后产生的是一个新slice，
			// 虽然父子调用关系上会可能会共享同一段底层数组（
			// append的属性，在cap足够时，append不会新开空间），
			// 但递归调用树同一层之间的局部变量path是彼此无关的，
			// 各个叶子节点的path之间也是完全不同的底层数组，
			// 因此不可能发生一个叶子的path被到达另一个叶子的过程中被修改的情况，
			// result保存时可以直接保存局部变量path。

			// dfs(nums, l, depth+1, append(path, num), visited, res)

			visited[i] = false // 每次返回后恢复状态，供下一回合使用
		}
	}
}

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

func permuteRecursion(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	var ret [][]int
	for i := 0; i < len(nums); i++ {
		buf := make([]int, len(nums)-1)
		copy(buf, nums[0:i])
		// fmt.Println(11, i, buf)

		copy(buf[i:], nums[i+1:])
		// fmt.Println(22, buf)

		r := permuteRecursion(buf)
		// fmt.Println(33, r)

		for _, v := range r {
			ret = append(ret, append(v, nums[i]))
		}
	}
	return ret
}

// @lc code=end

