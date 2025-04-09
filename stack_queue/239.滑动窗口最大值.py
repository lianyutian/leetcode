#
# @lc app=leetcode.cn id=239 lang=python3
#
# [239] 滑动窗口最大值
#

# @lc code=start
from collections import deque


class MyQueue:
    def __init__(self):
        self.queue = deque()  # 这里需要使用deque实现单调队列，直接使用list会超时

    def push(self, x):
        while self.queue and self.queue[-1] < x:
            self.queue.pop()
        self.queue.append(x)

    # 每次弹出的时候，比较当前要弹出的数值是否等于队列出口元素的数值，如果相等则弹出。
    # 同时pop之前判断队列当前是否为空。
    def pop(self, x):
        if self.queue and x == self.queue[0]:
            self.queue.popleft()  # list.pop()时间复杂度为O(n),这里需要使用collections.deque()

    # 查询当前队列里的最大值 直接返回队列前端也就是front就可以了。
    def front(self):
        return self.queue[0]


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        queue = MyQueue()
        res = []
        for i in range(k):  # 先将前k的元素放进队列
            queue.push(nums[i])
        res.append(queue.front())  # result 记录前k的元素的最大值
        for i in range(k, len(nums)):
            queue.pop(nums[i - k])  # 滑动窗口移除最前面元素
            queue.push(nums[i])  # 滑动窗口前加入最后面的元素
            res.append(queue.front())  # 记录对应的最大值
        return res


# @lc code=end
