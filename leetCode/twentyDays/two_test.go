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
