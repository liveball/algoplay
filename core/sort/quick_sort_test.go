package sort

import (
	"fmt"
	"testing"

	"github.com/liveball/algoplay/common"
	"github.com/liveball/algoplay/tools"
)

type args struct {
	list       common.List
	comparator common.Comparator
}

func TestQuickSort(t *testing.T) {

	// list1 := common.SliceToSimpleList(randomIntGenerator(10, 10))

	list1 := common.SliceToSimpleList(tools.New().Numbers(0, 99, 30))

	fmt.Println(list1)

	comparator1 := func(i, j int) bool {
		return list1.Get(i).(int) <= list1.Get(j).(int)
	}

	fmt.Println(list1)

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
			fmt.Println("before :", tt.args.list)
			QuickSort(tt.args.list, tt.args.comparator)

			fmt.Println("after :", tt.args.list)

			if !isSorted(list1, comparator1) {
				t.Fail()
			}
		})
	}
}

func TestQuickSortConcurrent(t *testing.T) {

	// list1 := common.SliceToSimpleList(randomIntGenerator(100000, 200000))
	list1 := common.SliceToSimpleList(tools.New().Numbers(0, 99, 10))

	comparator1 := func(i, j int) bool {
		return list1.Get(i).(int) <= list1.Get(j).(int)
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
			fmt.Println("before :", tt.args.list)

			QuickSortConcurrent(tt.args.list, tt.args.comparator)
			if !isSorted(list1, comparator1) {
				t.Fail()
			}

			fmt.Println("after :", tt.args.list)

		})
	}
}
