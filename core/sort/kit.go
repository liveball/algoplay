package sort

import (
	"github.com/liveball/algoplay/common"
	"math/rand"
	"time"
)

func exch(list common.List, i, j int) {
	old := list.Get(i)
	list.Set(i, list.Get(j))
	list.Set(j, old)
}

func judgeSorted(list common.List, comp common.Comparator) bool {
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
