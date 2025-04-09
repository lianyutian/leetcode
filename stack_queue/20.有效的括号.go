/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start

func isValid(s string) bool {
	stack := make([]rune, 0)

	m := make(map[rune]rune)
	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			// 拿到栈顶元素
			peek := stack[len(stack)-1]
			// 栈顶元素不匹配
			if peek != m[c] {
				return false
			}
			// 弹栈
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// @lc code=end
