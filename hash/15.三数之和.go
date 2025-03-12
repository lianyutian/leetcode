/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package hash

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {
	// [-4,-1,-1,0,1,2]
	sort.Ints(nums)
	l := len(nums)
	res := [][]int{}

	for left := 0; left < l-2; left++ {
		nl := nums[left]
		if nl > 0 {
			break
		}
		// 去重
		if left > 0 && nl == nums[left-1] {
			continue
		}

		index, right := left+1, l-1
		for index < right {
			ni, nr := nums[index], nums[right]
			sum := nl + ni + nr
			if sum == 0 {
				res = append(res, []int{nl, ni, nr})
				// 去重逻辑
				for index < right && nums[index] == ni {
					index++
				}
				for index < right && nums[right] == nr {
					right--
				}
			} else if sum < 0 {
				index++
			} else {
				right--
			}
		}
	}

	return res
}

// @lc code=end
