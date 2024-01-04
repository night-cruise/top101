package heapstackqueue

var m = map[byte]byte{'{': '}', '[': ']', '(': ')'}

func IsValid(s string) bool {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if m[s[i]] != 0 {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		bracket := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if m[bracket] != s[i] {
			return false
		}
	}

	return len(stack) == 0
}
