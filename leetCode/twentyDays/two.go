package twentyDays

/* 977. 有序数组的平方
https://leetcode-cn.com/problems/squares-of-a-sorted-array/submissions/
*/
func sortedSquares(nums []int) []int {
	reply := make([]int, len(nums))
	for k, v := range nums {
		nums[k] = v * v
	}

	head := 0
	tail := len(nums) - 1

	for index := len(nums) - 1; index >= 0; index-- {
		if nums[head] > nums[tail] {
			reply[index] = nums[head]
			head++
		} else {
			reply[index] = nums[tail]
			tail--
		}
	}
	return reply
}

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	strLen := len(nums)
	for i := 0; i < strLen/2; i++ {
		nums[i], nums[strLen-1-i] = nums[strLen-1-i], nums[i]
	}
}
