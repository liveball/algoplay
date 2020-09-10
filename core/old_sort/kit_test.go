package old_sort

import (
	"testing"

	"algoplay/common"
)

func TestIsSorted(t *testing.T) {
	type args struct {
		list common.List
		comp common.Comparator
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSorted(tt.args.list, tt.args.comp); got != tt.want {
				t.Errorf("isSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
