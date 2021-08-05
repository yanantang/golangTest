package twentyDays

/* 704. 二分查找
https://leetcode-cn.com/problems/binary-search/
*/

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	mid := -1
	for low < high {
		mid = (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if mid >= 0 && nums[mid] == target {
		return mid
	}
	return -1
}

/* 278. 第一个错误的版本
https://leetcode-cn.com/problems/first-bad-version/
*/

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	low := 1
	high := n
	mid := -1
	for low <= high {
		mid = (low + high) / 2
		if isBadVersion(mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

var bad = 4

func isBadVersion(version int) bool {
	return version >= bad
}

/* 35. 搜索插入位置
https://leetcode-cn.com/problems/search-insert-position/
*/

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	mid := -1
	if nums[high] < target {
		return high + 1
	}

	for low <= high {
		mid = (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return high + 1
}

// 题解优化后
func searchInsert1(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	mid := -1
	ans := len(nums)
	if nums[high] < target {
		return high + 1
	}

	for low <= high {
		mid = (low + high) / 2
		if nums[mid] >= target {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return ans
}
