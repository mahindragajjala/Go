package tailrecursion

func CountVowels(s string, index, count int) int {
	if index == len(s) {
		return count
	}
	c := s[index]
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
		count++
	}
	return CountVowels(s, index+1, count)
}
