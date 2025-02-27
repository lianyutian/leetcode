/*
 * @lc app=leetcode.cn id=242 lang=java
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
class Solution {
    public boolean isAnagram(String s, String t) {
        int[] record = new int[26];
        char[] sArr = s.toCharArray();
        for (int i = 0; i < sArr.length; i++) {
            record[sArr[i] - 'a']++;
        }
        char[] tArr = t.toCharArray();
        for (int i = 0; i < tArr.length; i++) {
            record[tArr[i] - 'a']--;
        }
        for (int i = 0; i < record.length; i++) {
            if (record[i] != 0) {
                return false;
            }
        }
        return true;
    }
}
// @lc code=end

