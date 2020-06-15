package main

import "fmt"

func main() {
	s := "()"
	fmt.Println(isValid(s))
	s = "()[]{}"
	fmt.Println(isValid(s))
}

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/
func isValid(s string) bool {
	// 空字符串，默认匹配
	if s == "" {
		return true
	}

	// 奇数，肯定不能匹配
	if len(s)%2 == 1 {
		return false
	}

	// 建立映射关系
	keyMap := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	// 堆栈
	var stack []string
	for i := 0; i < len(s); i++ {
		element := string(s[i])
		// 如果栈里有元素
		if len(stack) > 0 {
			// 如果和map里面的有匹配，也就是，目前元素是右半边括号
			if tmp, ok := keyMap[element]; ok {
				top := stack[len(stack)-1]
				// 【重要】如果当前遍历到的元素的映射和栈顶上的括号相同
				if top == tmp {
					// 栈顶元素pop掉
					stack = stack[:len(stack)-1]
					// 跳过这次循环
					continue
				}
			}
		}
		// 空栈或者没匹配上，就加入栈
		stack = append(stack, element)
	}

	// 如果最后栈是空的，说明都匹配上了，返回true
	return len(stack) == 0
}
