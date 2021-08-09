package twentyDays

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Printf("rotate %+v== []int{4,5,6,7,1,2,3}", nums)
}

func Test_moveZeroes(t *testing.T) {
	nums := []int{1, 0, 3, 0, 5, 6, 7}
	moveZeroes(nums)
	fmt.Printf("rotate %+v== []int{4,5,6,7,1,2,3}", nums)
}

func Test_twoSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	s := twoSum(nums, 7)
	fmt.Printf("rotate %+v== []int{4,5,6,7,1,2,3}", s)
}

func Test_reverseWords(t *testing.T) {
	s := reverseWords("Let's take LeetCode contest")
	fmt.Printf("reverseWords %+v== s'teL ekat edoCteeL tsetnoc", s)
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	s := lengthOfLongestSubstring("abcabcbb")
	fmt.Printf("lengthOfLongestSubstring %+v== 3", s)
}

func Test_floodFill(t *testing.T) {
	s := floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2)
	fmt.Printf("floodFill %+v== 3", s)
}

func Test_maxAreaOfIsland(t *testing.T) {
	s := maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	})
	fmt.Printf("maxAreaOfIsland %+v== 6", s)

	s = maxAreaOfIsland([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	})
	fmt.Printf("maxAreaOfIsland %+v== 4", s)

	s = maxAreaOfIsland([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	})
	fmt.Printf("maxAreaOfIsland %+v== 6", s)

}
