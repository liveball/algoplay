package sort

import "github.com/liveball/algoplay/common"

//ShellSort for shell sort
func ShellSort(l common.List, c common.Comparator, c2 common.Comparator) {
	n := l.Length()
	for step := n / 2; step > 0; step /= 2 {
		for i := step; i < n; i++ {
			if c(i, i-step) {

				k := i - step
				for k >= 0 && c2(k, i) {
					exchange(l, k+step, k)
					k -= step
				}

				exchange(l, k+step, i)
			}
		}
	}
}
