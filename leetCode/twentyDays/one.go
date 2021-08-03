package twentyDays


func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	mid := -1
	for low < high && high >= 0 {
		mid := (low + high) / 2
		if nums[mid] == target {
			break
		} else if nums[mid] < target {
			low = mid
		} else {
			high = mid
		}
	}

	if mid >= 0 && nums[mid] == target {
		return mid + 1
	}
	return -1
}
