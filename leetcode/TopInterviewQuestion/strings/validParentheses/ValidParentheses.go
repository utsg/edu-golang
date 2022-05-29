package leetcode

func IsValid(str string) bool {
	if len(str) % 2 != 0 {
		return false
	}

	parentheses := map[string]string{
		"(" : ")",
		"[" : "]",
		"{" : "}",
	}
	stack := []string{}

	for _, ch := range str {
		if _, ok := parentheses[string(ch)]; ok {
			stack = append(stack, string(ch))
		} else if len(stack) > 0 {
			// pop from stack
			popCh := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			if parentheses[popCh] != string(ch) {
				return false
			}
		} else {
			return false
		}
	}

	return len(stack) == 0
}
