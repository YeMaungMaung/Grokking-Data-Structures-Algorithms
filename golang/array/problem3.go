package array

import (
	"fmt"
	"math"
)

type Solution3 struct{}

func (s *Solution3) findDifferenceArray(nums []int) []int {
    n := len(nums)
    differenceArray := make([]int, n)
    leftSum := 0
    rightSum := 0

    for _, num := range nums {
        rightSum += num
    }

    for i := 0; i < n; i++ {
        rightSum -= nums[i]
        differenceArray[i] = int(math.Abs(float64(rightSum - leftSum)))
        leftSum += nums[i]
    }

    return differenceArray
}

func Test3() {
    solution := &Solution3{}

    example1 := []int{2, 5, 1, 6, 1}
    example2 := []int{3, 3, 3}
    example3 := []int{1, 2, 3, 4, 5}

    fmt.Println(solution.findDifferenceArray(example1)) // Output: [13 6 0 7 14]
    fmt.Println(solution.findDifferenceArray(example2)) // Output: [6 0 6]
    fmt.Println(solution.findDifferenceArray(example3)) // Output: [14 11 6 1 10]
}
