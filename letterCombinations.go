package main

import "fmt"

//题目：电话号码的字母组合
//leetcode链接：https://leetcode.cn/problems/letter-combinations-of-a-phone-number/

//根据题意，确定每个号码数字对应的字符，通过map来确定映射关系
var numToString map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var dst []string


//回溯法思路：输入的号码digits的每一位均需要与numToString进行映射，得到该位置的所有可能性，来与后续位进行组合，不断重复该逻辑，直到遍历完digits，得到每一条路径的组合
func getLetterCombinations(digits string)[]string{
	dst = []string{}

	//空输入""的情况
	if len(digits) == 0{
		return dst
	}


	backtrack(digits, 0, "")
	return dst
}

func backtrack(digits string, index int, combination string){
	if index == len(digits){
		dst = append(dst, combination)
	}else{
		num := string(digits[index])
		letters := numToString[num]
		for _, letter := range letters{
			backtrack(digits, index+1, combination+string(letter))
		}
	}
}


func main() {
	dights := "23"
	fmt.Println(getLetterCombinations(dights))
}
