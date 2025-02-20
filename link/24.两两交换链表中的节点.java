/*
 * @lc app=leetcode.cn id=24 lang=java
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode swapPairs(ListNode head) {
        ListNode dummyNode = new ListNode();
        dummyNode.next = head;
        ListNode pre = dummyNode;
        ListNode cur = head;

        while (cur != null && cur.next != null) {
            ListNode next = cur.next;
            ListNode nextnext = next.next;

            pre.next = next;
            next.next = cur;
            cur.next = nextnext;

            pre = cur;
            cur = nextnext;
        }

        return dummyNode.next;
    }
}
// @lc code=end

