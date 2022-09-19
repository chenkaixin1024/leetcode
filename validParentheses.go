package main

import "fmt"


//题目：有效的括号
//leetcode链接：https://leetcode.cn/problems/valid-parentheses/

//栈方法：利用栈的特性，左侧括号入栈，右侧括号与栈顶进行匹配，如果不匹配该字符串括号顺序不合法，
//而如果匹配，则继续直到遍历完字符串全部都能匹配上，且栈中没有字符剩余（栈顶=0）
func isVaild(s string) bool{

	l := len(s)
	//字符串长度为奇数，无法一一匹配，必然不合法
	if l % 2 != 0{
		return false
	}

	//左右括号匹配规则
	pattern := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, l)
	//栈顶
	idx := 0
	for _, c := range s{
		if c == '{' || c == '[' || c == '('{
			//入栈
			stack[idx] = byte(c)
			idx++
		}else{
			//出栈
			if idx == 0{
				return false
			}else if pattern[byte(c)] != stack[idx-1]{
				return false
			}else {
				idx--
			}
		}
	}

	if idx != 0{
		return false
	}

	return true
}

func main() {

	s := "[({()]"
	fmt.Println(isVaild(s))

}
