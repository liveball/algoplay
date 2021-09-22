package array

import (
	"fmt"
	"testing"
)

var score = []int{1, 1, -1, -1, -1, -1, 1}

func Test_qianzhuihe(t *testing.T) {
	fmt.Println("-----------")

	fmt.Println(10 & 7) //=2
	fmt.Println(^(2 ^ 7))

	fmt.Println("-----------")

	presum := make([]int, len(score)+1)
	for i := 0; i < len(score); i++ {
		presum[i+1] = score[i] + presum[i]
	}

	fmt.Println(presum)
	// [0 1 2 1 0 -1 -2 -1]
}

func TestSliceDel(t *testing.T) {
	sliceDel([]string{"a", "b", "c"}, "c")
	sliceDel2([]string{"a", "b", "c"}, "c")
	sliceDel3([]string{"a", "b", "c"}, "c")
}

// go test -run none -bench SliceDel -benchtime 3s -benchmem
// go test -run none -bench SliceDel -benchtime 3s -benchmem -o sliceDel -memprofile mem.out
// go tool pprof -alloc_space sliceDel mem.out
// list sliceDel
// go build -gcflags "-m -m"
func BenchmarkSliceDel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceDel([]string{"a", "b", "c"}, "c")
	}
}

// go test -run none -bench SliceDel2 -benchtime 3s -benchmem
// go test -run none -bench SliceDel2 -benchtime 3s -benchmem -o sliceDel2 -memprofile mem.out
// go tool pprof -alloc_space sliceDel2 mem.out
// list sliceDel2
// go build -gcflags "-m -m"

func BenchmarkSliceDel2(b *testing.B) {
	// fmt.Println(sliceDel2([]string{"a", "b", "c"}, "c"))
	for i := 0; i < b.N; i++ {
		sliceDel2([]string{"a", "b", "c"}, "c")
	}
}

func BenchmarkSliceDel3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceDel3([]string{"a", "b", "c"}, "c")
	}
}
