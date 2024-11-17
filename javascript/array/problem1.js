class Solution {
  runningSum(nums) {
    if (!nums || nums.length === 0) {
      return [];
    }

    const result = new Array(nums.length);

    result[0] = nums[0];
    for (let i = 1; i < nums.length; i++) {
      result[i] = result[i - 1] + nums[i];
    }

    return result;
  }
}

const solution = new Solution();

const testInputs = [
    [1,2,3,4],
    [3,1,4,2,2],
    [-1,-2,-3,-4,-5]
];

testInputs.forEach(input => {
    const output = solution.runningSum(input);
    
    console.log(output.join(' '));
});
