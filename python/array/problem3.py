class Solution:
    def findDifferenceArray(self, nums):
        n = len(nums)
        differenceArray = [0] * n
        leftSum = 0
        rightSum = sum(nums)

        for i in range(n):
            rightSum -= nums[i]
            differenceArray[i] = abs(rightSum - leftSum)
            leftSum += nums[i]

        return differenceArray

# Testing the solution
solution = Solution()

example1 = [2, 5, 1, 6, 1]
example2 = [3, 3, 3]
example3 = [1, 2, 3, 4, 5]

print(solution.findDifferenceArray(example1))  # Output: [13, 6, 0, 7, 14]
print(solution.findDifferenceArray(example2))  # Output: [6, 0, 6]
print(solution.findDifferenceArray(example3))  # Output: [14, 11, 6, 1, 10]
