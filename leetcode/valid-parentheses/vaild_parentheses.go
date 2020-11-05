package main

import "fmt"

var parenthesesMap = map[string]string{
	"[": "]",
	"{": "}",
	"(": ")",
}

type stack struct {
	values []string
}

func (s *stack) push(str string) {
	s.values = append(s.values, str)
}

func (s *stack) pop() string {
	length := len(s.values)
	if length < 1 {
		return ""
	}

	value := s.values[length-1]
	s.values = s.values[:length-1]
	return value
}

func (s *stack) len() int {
	return len(s.values)
}

func (s *stack) checkSatisfy() {
	if len(s.values) < 2 {
		return
	}

	right := s.pop()
	left := s.pop()
	if value, ok := parenthesesMap[left]; ok {
		if value == right {
			return
		}
	}

	s.push(left)
	s.push(right)
}

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	stk := new(stack)
	for i := 0; i < len(s); i++ {
		stk.push(fmt.Sprintf("%c", s[i]))
		stk.checkSatisfy()
	}

	if stk.len() == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
