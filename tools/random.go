package tools

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	Random struct {
	}
)

var (
	global = New()
)

func New() *Random {
	rand.Seed(time.Now().UnixNano())
	return new(Random)
}

func (r *Random) Strings() (s []string) {
	s = make([]string, 0)
	for i := 97; i < 123; i++ {
		s = append(s, fmt.Sprintf("%c ", rand.Intn(i)))
	}
	return
}

func (r *Random) Numbers(start int, end int, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}
	nums := make([]int, 0)
	for len(nums) < count {
		num := rand.Intn((end - start)) + start
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}
