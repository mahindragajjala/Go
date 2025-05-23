package tailrecursion

func CountChar(s string, ch byte, index, count int) int {
	if index == len(s) {
		return count
	}
	if s[index] == ch {
		count++
	}
	return CountChar(s, ch, index+1, count)
}
