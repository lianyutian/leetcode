/*
 * @lc app=leetcode.cn id=142 lang=java
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public ListNode detectCycle(ListNode head) {
        ListNode fast = head;
        ListNode slow = head;
        ListNode index = head;

        while (fast != null && fast.next != null) {
            fast = fast.next.next;
            slow = slow.next;

            // 存在环
            if (fast == slow) {
                // index 从头节点开始走，fast从环中开始走
                // 再次相遇时就是环入口
                while (fast != index) {
                    fast = fast.next;
                    index = index.next;
                }
                return index;
            }
        }
        return null;
    }
}
// @lc code=end

