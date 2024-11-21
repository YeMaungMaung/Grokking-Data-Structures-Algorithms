package array

import "fmt"

type Solution4 struct{}

func (s *Solution4) LargestAltitude(gain []int) int {
	currentAltitude := 0
	maxAltitude := 0

	for _, g := range gain {
		currentAltitude += g
		maxAltitude = max(currentAltitude, maxAltitude)
	}

	return maxAltitude
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test4() {
	solution := &Solution4{}
	fmt.Println(solution.LargestAltitude([]int{-5, 1, 5, 0, -7})) // Expected: 1
	fmt.Println(solution.LargestAltitude([]int{4, -3, 2, -1, -2})) // Expected: 4
	fmt.Println(solution.LargestAltitude([]int{2, 2, -3, -1, 2, 1, -5})) // Expected: 4
}
