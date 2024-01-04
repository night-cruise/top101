package heapstackqueue

// 双栈
func Solve(s string) int {
	var opPriority = map[byte]int{
		'(': 3,
		'+': 1,
		'-': 1,
		'*': 2,
	}
	stack1, stack2 := []int{}, []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack2 = append(stack2, s[i])
		} else if s[i] == '+' || s[i] == '*' || s[i] == '-' {
			for len(stack2) > 0 && stack2[len(stack2)-1] != '(' && opPriority[stack2[len(stack2)-1]] >= opPriority[s[i]] {
				op := stack2[len(stack2)-1]
				stack2 = stack2[:len(stack2)-1]
				num1, num2 := stack1[len(stack1)-2], stack1[len(stack1)-1]
				stack1 = stack1[:len(stack1)-2]
				switch op {
				case '+':
					stack1 = append(stack1, num1+num2)
				case '-':
					stack1 = append(stack1, num1-num2)
				case '*':
					stack1 = append(stack1, num1*num2)
				}
			}
			stack2 = append(stack2, s[i])
		} else if s[i] == ')' {
			for stack2[len(stack2)-1] != '(' {
				op := stack2[len(stack2)-1]
				stack2 = stack2[:len(stack2)-1]
				num1, num2 := stack1[len(stack1)-2], stack1[len(stack1)-1]
				stack1 = stack1[:len(stack1)-2]
				switch op {
				case '+':
					stack1 = append(stack1, num1+num2)
				case '-':
					stack1 = append(stack1, num1-num2)
				case '*':
					stack1 = append(stack1, num1*num2)
				}
			}
			stack2 = stack2[:len(stack2)-1]
		} else {
			num := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			stack1 = append(stack1, num)
			i--
		}
	}
	for len(stack1) > 1 {
		op := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		num1, num2 := stack1[len(stack1)-2], stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-2]
		switch op {
		case '+':
			stack1 = append(stack1, num1+num2)
		case '-':
			stack1 = append(stack1, num1-num2)
		case '*':
			stack1 = append(stack1, num1*num2)
		}
	}

	return stack1[0]
}
