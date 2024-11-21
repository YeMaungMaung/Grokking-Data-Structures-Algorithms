class Solution:
    def largestAltitude(self, gain):
        current_altitude = 0
        max_altitude = 0

        for i in gain:
            current_altitude += i
            max_altitude = max(current_altitude, max_altitude)

        return max_altitude

if __name__ == "__main__":
    solution = Solution()
    print(solution.largestAltitude([-5, 1, 5, 0, -7]))  # Expected: 1
    print(solution.largestAltitude([4, -3, 2, -1, -2]))  # Expected: 4
    print(solution.largestAltitude([2, 2, -3, -1, 2, 1, -5]))  # Expected: 4
