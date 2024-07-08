package stack

import (
	"strconv"
)

//逆波兰式表达式求值，中文版LeetCode，第150题

/*
算法思路：

  - 遍历字符串数组，遇到非“+-/*”，那么其就是字符串形式的数字，将其转为数字，入栈。遇到符号，判断其具体是哪个符号。

  - 第二步，把栈内的最后两个元素出栈，并计算表达式的结果，把这个结果压回栈内。

  - 第三步，重复第一步。

  - 最后栈内只会剩下一个数字，返回即可

    运用到的包strconv,一个可以将字符转为数字的包
*/
func EvalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil { //遇到数字，入栈
			stack = append(stack, val)
		} else {
			// 遇到，"+_?/"
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2] //去掉已经计算的两个，然后把计算结果压栈
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			default:
				stack = append(stack, num1/num2)

			}
		}
	}
	return stack[0]
}
