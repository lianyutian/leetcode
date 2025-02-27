#
# @lc app=leetcode.cn id=242 lang=python3
#
# [242] 有效的字母异位词
#

# @lc code=start
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        record = [0] * 26
        for c in s:
            # ord() 函数用于将字符转换为相应的整数值。
            record[ord(c) - ord('a')] += 1
        for c in t:
            record[ord(c) - ord('a')] -= 1

        for v in record:
            if v != 0:
                return False
            

        return True

        
# @lc code=end

