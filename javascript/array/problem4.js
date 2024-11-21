class Solution {
  largestAltitude(gain) {
    let currentAltitude = 0;
    let maxAltitude = 0;

    for (let i of gain) {
      currentAltitude += i;
      maxAltitude = Math.max(currentAltitude, maxAltitude);
    }

    return maxAltitude;
  }
}

// Examples
const solution = new Solution();
console.log(solution.largestAltitude([-5, 1, 5, 0, -7])); // Expected: 1
console.log(solution.largestAltitude([4, -3, 2, -1, -2])); // Expected: 4
console.log(solution.largestAltitude([2, 2, -3, -1, 2, 1, -5])); // Expected: 4
