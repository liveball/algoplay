package main

func main() {

}

func getSecMax(nums []int) int {
	max := 0
	secMax := 0
	for _, num := range nums {
		if num > max {
			secMax = max
			max = num
			continue
		}
		if num > secMax {
			secMax = num
		}
	}

	return secMax
}
