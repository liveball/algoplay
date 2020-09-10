package main

import "fmt"

func main() {
	GuessingGame()
}

// As a more whimsical example, this program guesses your number:
//
func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")

	answer := Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)

		fmt.Scanf("%s", &s)

		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}
