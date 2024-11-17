package array

import "fmt"

type Solution struct{}

func (s *Solution) runningSum(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }

    result := make([]int, len(nums))
    result[0] = nums[0]

    for i := 1; i < len(nums); i++ {
        result[i] = result[i-1] + nums[i]
    }

    return result
}

func Test1() {
    solution := &Solution{}

    // Test cases
    testInputs := [][]int{
        {1, 2, 3, 4},
        {3, 1, 4, 2, 2},
        {-1, -2, -3, -4, -5},
    }

    for _, input := range testInputs {
        output := solution.runningSum(input)

        // Print the output array
        for _, val := range output {
            fmt.Print(val, " ")
        }
        fmt.Println()
    }
}
