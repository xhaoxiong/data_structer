package main

import (
	"data_structer/stack"
	"github.com/spf13/cast"
)

func main() {
}

func isValid(s string) bool {


	stack := stack.InitNoParam()
	for _, ss := range s {

		char := string(ss)
		if char == "{" || char == "[" || char == "(" {
			stack.Push(char)
		} else {
			if stack.IsEmpty() {
				return false
			}

			if cast.ToString(stack.Peek()) == "(" && char != ")" {
				return false
			}
			if cast.ToString(stack.Peek()) == "[" && char != "]" {
				return false
			}
			if cast.ToString(stack.Peek()) == "{" && char != "}" {
				return false
			}
			stack.Pop()
		}
	}
	return stack.IsEmpty()
}
