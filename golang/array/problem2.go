package array

import "fmt"

type Solution2 struct{}

func (s *Solution2) containsDuplicate(nums []int) bool {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                return true
            }
        }
    }
    return false
}

func Test2() {
    solution := Solution2{}

    nums1 := []int{1, 2, 3, 4}
    fmt.Println(solution.containsDuplicate(nums1)) // Expected output: false

    nums2 := []int{1, 2, 3, 1}
    fmt.Println(solution.containsDuplicate(nums2)) // Expected output: true

    nums3 := []int{}
    fmt.Println(solution.containsDuplicate(nums3)) // Expected output: false

    nums4 := []int{1, 1, 1, 1}
    fmt.Println(solution.containsDuplicate(nums4)) // Expected output: true
}
