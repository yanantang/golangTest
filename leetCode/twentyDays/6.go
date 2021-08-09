package twentyDays
/*3. 无重复字符的最长子串
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 */
func lengthOfLongestSubstring(s string) int {
	ret, i := 0, 0

	for  i < len(s) {
		tmp := 0
		strMap := make(map[byte]int)
		for ; i < len(s); i++ {
			if index, ok := strMap[s[i]]; ok {
				delete(strMap, s[i])
				i = index + 1
				break
			}
			tmp++
			strMap[s[i]] = i
		}

		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}
