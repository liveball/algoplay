package array

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)

	res := [][]int{}
	zeroNum := 0

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				fnum := nums[i]
				snum := nums[j]
				tnum := nums[k]
				rows := []int{fnum, snum, tnum}

				if fnum+snum+tnum == 0 {
					if fnum == 0 && snum == 0 && tnum == 0 && zeroNum == 0 {
						res = append(res, rows)
						zeroNum++
						fmt.Println("zeroNum repeat", zeroNum)
						continue
					}

					if !isContains(res, rows) {
						res = append(res, rows)
					}
				}
			}
		}
	}

	return res
}

func isContains(fList [][]int, iList []int) bool {
	for _, list := range fList {
		if testEq(list, iList) {
			return true
		}
	}
	return false
}

func testEq(a, b []int) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func isExist(t []int, arr [][]int) bool {
	rows := len(arr)
	columns := len(arr[0])

	row, column := 0, 0

	cnt := 0
	if rows > 0 && columns > 0 {
		for row < rows && column < columns {
			if arr[row][column] == 0 {
				cnt++
			}

			row++
			column++
		}
	}

	fmt.Println(cnt)

	return false
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				t := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = t
			}
		}
	}
}
