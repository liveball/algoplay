package main

import (
	"fmt"
	"sync"

	"github.com/liveball/algoplay/common"
	"github.com/liveball/algoplay/core/sort"
	"github.com/liveball/algoplay/tools"
)

//GOMAXPROCS=1
//go build -o main main.go && GODEBUG=schedtrace=1000,scheddetail=1 ./main

func main() {

	pool := common.New(2)

	var wg sync.WaitGroup
	wg.Add(1)
	pool.Go(func() {
		fmt.Println("hello alg")
		wg.Done()
	})
	wg.Wait()
	pool.Close() //让工作池停止工作， 等待所有现有的工作完成

	testQuickSortConcurrent()
	// time.Sleep(1 * time.Minute)
}

type args struct {
	list       common.List
	comparator common.Comparator
}

func testQuickSortConcurrent() {

	// list1 := common.SliceToSimpleList(randomIntGenerator(100000, 200000))
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
		fmt.Println("before :", tt.args.list)

		sort.QuickSortConcurrent(tt.args.list, tt.args.comparator)

		fmt.Println("after :", tt.args.list)

	}
}
