package old_sort

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"algoplay/common"
)

func exchange(list common.List, i, j int) {
	old := list.Get(i)
	list.Set(i, list.Get(j))
	list.Set(j, old)
}

func isSorted(list common.List, comp common.Comparator) bool {
	for i := 0; i < list.Length()-1; i++ {
		if !comp(i, i+1) {
			return false
		}
	}
	return true
}

func randomIntGenerator(count int, max int) []int {
	rand.Seed(time.Now().Unix())
	res := make([]int, count)
	for i := 0; i < count; i++ {
		res[i] = int(rand.Float32() * float32(max))
	}
	return res
}

func listToString(list common.List) string {
	builder := &strings.Builder{}
	builder.WriteString("[ ")
	for i, v := range list.ToSlice() {
		builder.WriteString(fmt.Sprintf("%d: %v, ", i, v))
	}
	builder.WriteString(" ]")
	return builder.String()
}
