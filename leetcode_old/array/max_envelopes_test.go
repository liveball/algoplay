package array

import (
	"fmt"
	"testing"
)

func Test_maxEnvelopes(t *testing.T) {
	envelopes := [][]int{
		{5,4},
		{6,4},
		{6,7},
		{2,3},
	}

	res:=maxEnvelopes(envelopes)
	fmt.Println(res)
}
