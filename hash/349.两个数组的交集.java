/*
 * @lc app=leetcode.cn id=349 lang=java
 *
 * [349] 两个数组的交集
 */

// @lc code=start
class Solution {
    public int[] intersection(int[] nums1, int[] nums2) {
        Set<Integer> tempMap = new HashSet<>();
        
        int k = 0;

        for (int num : nums1) {
            if (!tempMap.contains(num)) {
                tempMap.add(num);
            }
        }
        List<Integer> resList = new ArrayList<>();
        for (int num : nums2) {
            if (tempMap.contains(num)) {
                resList.add(num);
                tempMap.remove(num);
            }
        }

        int[] res = new int[resList.size()];

        for (int i = 0; i < resList.size(); i++) {
            res[i] = resList.get(i);
        }

        return res;
    }
}
// @lc code=end

