# @before-stub-for-debug-begin
from python3problem225 import *
from typing import *

# @before-stub-for-debug-end

#
# @lc app=leetcode.cn id=225 lang=python3
#
# [225] 用队列实现栈
#
from collections import deque


# @lc code=start
class MyStack:

    def __init__(self):
        self.queue = deque()

    def push(self, x: int) -> None:
        self.queue.append(x)

    def pop(self) -> int:
        if self.empty():
            return None
        return self.queue.pop()

    def top(self) -> int:
        if self.empty():
            return None
        val = self.pop()
        self.queue.append(val)
        return val

    def empty(self) -> bool:
        return not self.queue


# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()
# @lc code=end
