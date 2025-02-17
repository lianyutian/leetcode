/*
 * @lc app=leetcode.cn id=59 lang=java
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
class Solution {
    public int[][] generateMatrix(int n) {
        int top = 0;
        int right = n - 1;
        int bottom = n - 1;
        int left = 0;
        int k = 1;
        int[][] res = new int[n][n];

        while (k < (n * n)) {
            // 从左往右，找一条不变的边界（上边界）
            for (int i = left; i < right; i++) {
                res[top][i] = k;
                k++;
            }
            // 从上往下，找一条不变的边界（右边界）
            for (int i = top; i < bottom; i++) {
                res[i][right] = k;
                k++;
            }
            // 从右往左，找一条不变的边界（下边界）
            for (int i = right; i > left; i--) {
                res[bottom][i] = k;
                k++;
            }
            // 从下往上，找一条不变的边界（左边界）
            for (int i = bottom; i > top; i--) {
                res[i][left]= k;
                k++;
            }
            // 绕一圈后边界往里缩
            top++;
            right--;
            bottom--;
            left++;
        }

        // 中心点
        res[left][right] = n * n;

        return res;
    }
}
// @lc code=end

