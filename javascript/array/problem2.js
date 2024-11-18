class Solution {
  containsDuplicate(nums) {
    for (let i = 0; i < nums.length; i++) {
      for (let j = i + 1; j < nums.length; j++) {
        if (nums[i] === nums[j]) {
          return true;
        }
      }
    }
    return false;
  }
}

const sol = new Solution();
const nums1 = [1, 2, 3, 4];
console.log(sol.containsDuplicate(nums1)); // Expected output: false

const nums2 = [1, 2, 3, 1];
console.log(sol.containsDuplicate(nums2)); // Expected output: true

const nums3 = [];
console.log(sol.containsDuplicate(nums3)); // Expected output: false

const nums4 = [1, 1, 1, 1];
console.log(sol.containsDuplicate(nums4)); // Expected output: true
