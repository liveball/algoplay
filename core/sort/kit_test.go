package sort

import (
	"testing"

	"github.com/liveball/algoplay/common"
)

func Test_judgeSorted(t *testing.T) {
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
			if got := judgeSorted(tt.args.list, tt.args.comp); got != tt.want {
				t.Errorf("judgeSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
