package twentyDays

/* 344. 反转字符串
https://leetcode-cn.com/problems/reverse-string/submissions/
*/
func reverseString(s []byte) {
	strLen := len(s)
	for i := 0; i < strLen/2; i++ {
		s[i], s[strLen-i-1] = s[strLen-i-1], s[i]
	}
}

/* 557. 反转字符串中的单词 III
https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/submissions/
 */
func reverseWords(s string) string {
	arrs := []byte(s)
	space := 0
	for i, v := range arrs {
		if v != ' ' {
			continue
		}

		for j := space; j < (i+space)/2; j++ {
			arrs[j], arrs[i-j-1+space] = arrs[i-j-1+space], arrs[j]
		}
		space = i + 1
	}

	for j := space; j < (len(s)+space)/2; j++ {
		arrs[j], arrs[len(s)-j-1+space] = arrs[len(s)-j-1+space], arrs[j]
	}
	return string(arrs)
}
