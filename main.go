package main

import (
	"fmt"
	"runtime"

	"algoplay/common"
	"algoplay/core/sort"
	"algoplay/tools"
)

//GOMAXPROCS=1
//go build -o main main.go && GODEBUG=schedtrace=1000,scheddetail=1 ./main

func main() {
	println("start NumGoroutine:", runtime.NumGoroutine())

	testQuickSort()
	sort.Close()
	println("end NumGoroutine:", runtime.NumGoroutine())
}

type args struct {
	list       common.List
	comparator common.Comparator
}

func testQuickSort() {
	// list1 := common.SliceToSimpleList(tools.New().Numbers(0, 20, 20))
	list1 := common.SliceToSimpleList(tools.New().Numbers(0, 200000, 100000))

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
		// fmt.Println("before :", tt.args.list)

		// sort.QuickSort(tt.args.list, tt.args.comparator)
		sort.QuickSortConcurrent(tt.args.list, tt.args.comparator)

		fmt.Println("after :", tt.args.list)

	}
}
