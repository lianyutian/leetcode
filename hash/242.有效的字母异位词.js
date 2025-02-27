/*
 * @lc app=leetcode.cn id=242 lang=javascript
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
    const record = new Array(26).fill(0);

    const base = "a".charCodeAt();
    for (let i = 0; i < s.length; i++) {
        record[s.charCodeAt(i) - base]++
    }

    for (let i = 0; i < t.length; i++) {
        record[t.charCodeAt(i) - base]--
    }

    for (let i = 0; i < record.length; i++) {
        if (record[i] != 0) {
            return false
        }
    }

    return true
};
// @lc code=end

