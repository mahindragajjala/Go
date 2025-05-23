package tailrecursion

func ReverseStringTail(s string, acc string) string {
	if len(s) == 0 {
		return acc
	}
	return ReverseStringTail(s[1:], string(s[0])+acc)
}
