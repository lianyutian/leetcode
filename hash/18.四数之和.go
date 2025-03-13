/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	var res [][]int

	for left := 0; left < l-3; left++ {
		nl := nums[left]
		// 去重
		if left > 0 && nums[left-1] == nl {
			continue
		}

		for left2 := left + 1; left2 < l-2; left2++ {
			nl2 := nums[left2]

			index, right := left2+1, l-1
			// 去重
			if left2 > left+1 && nums[left2-1] == nl2 {
				continue
			}
			for index < right {
				ni, nr := nums[index], nums[right]

				sum := nl + nl2 + ni + nr

				if sum == target {
					res = append(res, []int{nl, nl2, ni, nr})
					// 去重
					for index < right && nums[index] == ni {
						index++
					}
					for index < right && nums[right] == nr {
						right--
					}
				} else if sum < target {
					index++
				} else {
					right--
				}
			}
		}
	}
	return res
}

// @lc code=end
