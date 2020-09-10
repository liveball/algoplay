package ten_sort

import "testing"

func Test_insertSort(t *testing.T){
	nums:=[]int{6,7,4,5,2,3,1}

	insertSort(nums)

	t.Log(nums)
}
