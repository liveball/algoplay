package sort

import (
	"testing"

	"github.com/liveball/algoplay/common"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		list       common.List
		comparator common.Comparator
	}
	list1 := common.SliceToSimpleList([]common.Any{3, 2, 6, 4, 1})
	comparator1 := func(i, j int) bool {
		return list1.Get(i).(int) < list1.Get(j).(int)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				list:       list1,
				comparator: comparator1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.list, tt.args.comparator)
			if !judgeSorted(list1, comparator1) {
				t.Fail()
			}
		})
	}
}