class Solution {
  findDifferenceArray(nums) {
    const n = nums.length;
    const differenceArray = new Array(n);
    let leftSum = 0;
    let rightSum = nums.reduce((acc, num) => acc + num, 0);

    for (let i = 0; i < n; i++) {
      rightSum -= nums[i];
      differenceArray[i] = Math.abs(rightSum - leftSum);
      leftSum += nums[i];
    }

    return differenceArray;
  }
}

// Testing the solution
const solution = new Solution();

const example1 = [2, 5, 1, 6, 1];
const example2 = [3, 3, 3];
const example3 = [1, 2, 3, 4, 5];

console.log(solution.findDifferenceArray(example1)); // Output: [13, 6, 0, 7, 14]
console.log(solution.findDifferenceArray(example2)); // Output: [6, 0, 6]
console.log(solution.findDifferenceArray(example3)); // Output: [14, 11, 6, 1, 10]
