/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
// @lc code=start
type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (q *MyQueue) Pop(x int) {
	if len(q.queue) > 0 && q.queue[0] == x {
		q.queue = q.queue[1:]
	}
}

func (q *MyQueue) Push(x int) {
	for len(q.queue) > 0 && q.queue[len(q.queue)-1] < x {
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue, x)
}

func (q *MyQueue) Front() int {
	return q.queue[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	MyQueue := NewMyQueue()
	for i := 0; i < k; i++ {
		MyQueue.Push(nums[i])
	}
	res := []int{MyQueue.Front()}
	for i := k; i < len(nums); i++ {
		// 右移
		MyQueue.Pop(nums[i-k])
		MyQueue.Push(nums[i])
		res = append(res, MyQueue.Front())
	}
	return res
}

// @lc code=end
