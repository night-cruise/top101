package recursionbacktracking

func GenerateParenthesis(n int) []string {
	brackets, ans := make([]byte, 2*n), []string{}

	GPHelper(brackets, 0, &ans, 0, 0)

	return ans
}
func GPHelper(brackets []byte, i int, ans *[]string, leftbracketCnt int, rightbracketCnt int) {
	if i == len(brackets) {
		*ans = append(*ans, string(brackets))
		return
	}
	if leftbracketCnt == len(brackets)/2 {
		brackets[i] = ')'
		GPHelper(brackets, i+1, ans, leftbracketCnt, rightbracketCnt+1)
	} else if leftbracketCnt == rightbracketCnt {
		brackets[i] = '('
		GPHelper(brackets, i+1, ans, leftbracketCnt+1, rightbracketCnt)
	} else {
		brackets[i] = '('
		GPHelper(brackets, i+1, ans, leftbracketCnt+1, rightbracketCnt)
		brackets[i] = ')'
		GPHelper(brackets, i+1, ans, leftbracketCnt, rightbracketCnt+1)
	}
}
