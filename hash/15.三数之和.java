/*
 * @lc app=leetcode.cn id=15 lang=java
 *
 * [15] 三数之和
 */

// @lc code=start
class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        int l = nums.length;
        List<List<Integer>> res = new ArrayList<>();
        for (int left = 0; left < l - 2; left++) {
            int nl = nums[left];
            if (nl > 0) {
                break;
            }
            if (left > 0 && nl == nums[left - 1]) {
                continue;
            }
            int index = left + 1;
            int right = l - 1;
            while (index < right) {
                int ni = nums[index];
                int nr = nums[right];
                int sum = nl + ni + nr;
                
                if (sum == 0) {
                    List<Integer> list = new ArrayList<>();
                    list.add(nl);
                    list.add(ni);
                    list.add(nr);
                    res.add(list);

                    while (index < right && ni == nums[index]) {
                        index++;
                    }
                    while (index < right && nr == nums[right]) {
                        right--;
                    }
                } else if (sum < 0) {
                    index++;
                } else {
                    right--;
                }
            }
        }
        return res;
    }
}
// @lc code=end

