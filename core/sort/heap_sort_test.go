package sort

import (
	"testing"

	"algoplay/common"
)

func Test_parent(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 0},
		{"2", args{2}, 0},
		{"3", args{3}, 1},
		{"4", args{4}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parent(tt.args.i); got != tt.want {
				t.Errorf("parent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildMaxHeap(t *testing.T) {
	type args struct {
		list common.List
		comp common.Comparator
	}
	list1 := common.SliceToSimpleList([]int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7})
	comp1 := func(i, j int) bool {
		return list1.Get(i).(int) <= list1.Get(j).(int)
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{list1, comp1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BuildMaxHeap(tt.args.list, tt.args.comp)
			if !IsMaxHeap(tt.args.list, tt.args.comp) {
				t.Fail()
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	type args struct {
		list common.List
		comp common.Comparator
	}
	list1 := common.SliceToSimpleList(randomIntGenerator(1000000, 1000000))
	comp1 := func(i, j int) bool {
		return list1.Get(i).(int) <= list1.Get(j).(int)
	}

	tests := []struct {
		name string
		args args
	}{
		{"1", args{list1, comp1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.list, tt.args.comp)
			if !isSorted(tt.args.list, tt.args.comp) {
				t.Fail()
			}
		})
	}
}
