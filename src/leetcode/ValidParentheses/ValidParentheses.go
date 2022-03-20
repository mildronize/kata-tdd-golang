/*
Problem: https://leetcode.com/problems/valid-parentheses/
Author: Thada Wangthammang
*/
package ValidParentheses

import "github.com/emirpasic/gods/stacks/arraystack"

func isValid(s string) bool {
	stack := arraystack.New()
	for _, ch := range s {
		if ch == '{' || ch == '(' {
			stack.Push(ch)
		} else if ch == '}' || ch == ')' {
			peekValue, _ := stack.Peek()
			if peekValue == '{' {
				stack.Pop()
			} else if peekValue == '(' {
				stack.Pop()
			}
		}
	}
	return stack.Size() == 0
}
