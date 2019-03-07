package sort

import (
	"fmt"
	"testing"

	"github.com/liveball/algoplay/common"
	"github.com/liveball/algoplay/tools"
)

type shellargs struct {
	list                    common.List
	comparator, comparator2 common.Comparator
}

func TestShellSort(t *testing.T) {

	list1 := common.SliceToSimpleList(tools.New().Numbers(0, 100, 50))

	cp1 := func(i, j int) bool {
		return list1.Get(i).(int) < list1.Get(j).(int)
	}

	cp2 := func(i, j int) bool {
		return list1.Get(i).(int) > list1.Get(j).(int)
	}

	tests := []struct {
		name string
		args shellargs
	}{
		{
			name: "1",
			args: shellargs{
				list:        list1,
				comparator:  cp1,
				comparator2: cp2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("before :", tt.args.list)

			ShellSort(tt.args.list, tt.args.comparator, tt.args.comparator2)

			fmt.Println("after :", tt.args.list)
			if !isSorted(list1, cp1) {
				t.Fail()
			}
		})
	}
}
