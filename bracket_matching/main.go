package main

import (
	"important_test/data_structer/stack"
	"github.com/spf13/cast"
	"fmt"
)

func main() {
	fmt.Println(isValid())
}

func isValid() bool {
	s := "{[(])}"

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
