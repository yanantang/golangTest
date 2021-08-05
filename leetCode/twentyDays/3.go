package twentyDays

/* 283. 移动零
https://leetcode-cn.com/problems/move-zeroes/submissions/
*/
func moveZeroes(nums []int) {
	zp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && nums[zp] == 0 {
			swap(nums, i, zp)
			zp++
		} else if nums[i] == 0 && nums[zp] != 0 {
			zp = i
		}
	}
}

func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
/* 167. 两数之和 II - 输入有序数组
https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
 */
func twoSum(numbers []int, target int) []int {
	reply := make([]int, 2)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers) && i != j && numbers[i]+numbers[j] <= target; j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i, j}
			}
		}
	}
	return reply
}
