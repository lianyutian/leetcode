/*
 * @lc app=leetcode.cn id=19 lang=javascript
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} n
 * @return {ListNode}
 */
var removeNthFromEnd = function(head, n) {
    let dummyNode = new ListNode()
    dummyNode.next = head
    let fast = head
    let cur = dummyNode

    for (let i = 0; i < n; i++) {
        fast = fast.next
    }

    while (fast != null) {
        fast = fast.next
        cur = cur.next
    }

    cur.next = cur.next.next

    return dummyNode.next
};
// @lc code=end

