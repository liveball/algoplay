package ten_sort

import "testing"

func Test_selectSort(t *testing.T){
	nums:=[]int{6,7,4,5,2,3,1}

	selectSort(nums)

	t.Log(nums)
}
