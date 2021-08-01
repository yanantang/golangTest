package week252

/*
地址: https://leetcode-cn.com/contest/weekly-contest-252/
完成:1,2,3, 4(未做)
*/
import (
	"sort"
	"strconv"
	"strings"
)

func isThree(n int) bool {
	var count int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
		if count > 3 {
			return false
		}
	}
	return count == 3
}

func numberOfWeeks(milestones []int) int64 {
	if len(milestones) == 1 {
		return 1
	}

	sort.Ints(milestones)
	diff := milestones[len(milestones)-1]

	var total int64
	for i := 0; i < len(milestones); i++ {
		if i != len(milestones)-1 {
			diff -= milestones[i]
		}
		total += int64(milestones[i])
	}
	if diff > 0 {
		total -= int64(diff)
	}

	return total
}

func minimumPerimeter(neededApples int64) int64 {
	var depth int64 = 1
	for ; neededApples > 0; depth++ {
		var sum int64
		if depth == 1 {
			sum = 3
		} else {
			sum = (depth+1)*(depth+2*depth) - depth*3
		}

		neededApples -= sum * 4
	}

	if depth == 1 {
		return 8
	}
	return (depth - 1) * 8
}

func largestOddNumber(num string) string {
	arr := []byte(num)

	var i = len(arr) - 1
	for ; i >= 0; i-- {
		if arr[i]%2 == 1 {
			break
		}
	}

	if i >= 0 {
		return string(arr[0 : i+1])
	}
	return ""
}

func numberOfRounds(startTime string, finishTime string) int {
	startArr := strings.Split(startTime, ":")
	finishArr := strings.Split(finishTime, ":")

	times := map[int]int{
		0:  0,
		15: 1,
		30: 2,
		45: 3,
	}

	s0, _ := strconv.Atoi(startArr[0])
	s1, _ := strconv.Atoi(startArr[1])

	f0, _ := strconv.Atoi(finishArr[0])
	f1, _ := strconv.Atoi(finishArr[1])

	sSum := s0 * 4
	if t, ok := times[s1]; ok {
		sSum += t
	} else {
		sSum = sSum + s1/15 + 1
	}

	fSum := f0 * 4
	if t, ok := times[f1]; ok {
		fSum += t
	} else {
		fSum = fSum + f1/15
	}

	if fSum-sSum < 0 && startTime > finishTime {
		return fSum - sSum + 24*4
	} else if fSum-sSum > 0 {
		return fSum - sSum
	}
	return 0
}
