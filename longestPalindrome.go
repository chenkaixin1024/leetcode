package main

import "fmt"


//动态规划-以map[string]bool存储回文子串，截取子串操作太多，性能较低（应减少子串的截取）
func longestPalindromeWay1(s string) string{

	strlen := len(s)
	if strlen < 2{
		return s
	}

	pailinerome := make(map[string]bool)
	dst := s[0:1]

	for _, sub := range s{
		pailinerome[string(sub)] = true
	}

	for l := 2; l <= strlen; l++{

		for i := 0; i < strlen; i++{

			j := i + l
			if j > strlen{
				break
			}
			if s[i] == s[j-1]{
				if l == 2{
					pailinerome[s[i:j]] = true
				}else {
					pailinerome[s[i:j]] = pailinerome[s[i+1:j-1]]
				}
			}else{
				pailinerome[s[i:j]] = false
			}

			if pailinerome[s[i:j]] && l > len(dst){
				dst = s[i:j]
			}

		}
	}

	//fmt.Println(pailinerome)

	return dst

}


//相较way1，同样是动态规划，对比不同的是利用二维数组存回文子串的左右边界，得到最长的回文子串，在返回值中截取返回
//动态规划思路：
//s[i:j]为回文子串的前提(1.s[i]==s[j]；2.当字符串长度大于2时，s[i+1][j-1]为回文子串；小于2时，满足条件1即可)
func longestPalindromeWay2(s string)string{
	strlen := len(s)
	//字符串长度小于2，本身就是回文
	if strlen < 2{
		return s
	}

	//二维数组-用于存放回文子串的左右边界
	pailinerome := make([][]bool, strlen)
	for k := range pailinerome{
		pailinerome[k] = make([]bool, strlen)
		//子串长度=1，均是回文子串
		pailinerome[k][k] = true
	}
	left := 0
	max := 1

	//l为子串长度
	for l := 2; l <= strlen; l++{

		for i := 0; i < strlen; i++{

			j := i + l - 1
			if j >= strlen{
				break
			}
			if s[i] == s[j]{
				if l == 2{
					pailinerome[i][j] = true
				}else {
					pailinerome[i][j] = pailinerome[i+1][j-1]
				}
			}

			//当前子串为回文且长度大于目前的最长子串时，替换成为最长子串
			if pailinerome[i][j] && l > max{
				left = i
				max = l
			}
		}
	}

	//fmt.Println(pailinerome)

	return s[left:left+max]
}


//中心扩展法思路：
//回文子串分为奇数长度子串和偶数长度子串，均满足由中心向两边扩展时，左右边界相等
//以此为条件，验证得到字符串不同中心的最长回文子串
func longestPalindromeWay3(s string)string{

	strlen := len(s)
	if strlen < 2{
		return s
	}


	max := 1
	left := 0
	right := 0
	for i := 0; i < strlen; i++{
		//长度为奇数的子串
		left1, right1 := expandAroundCenter(s,i,i)
		//长度为偶数的子串
		left2, right2 := expandAroundCenter(s,i,i+1)

		if right1 - left1 + 1 > max{
			left = left1
			right = right1
			max = right - left + 1
		}

		if right2 - left2 + 1 > max{
			left = left2
			right = right2
			max = right - left + 1
		}

	}


	return s[left:right+1]
}


//由中心向两边递推，验证是否满足回文条件s[left]==s[right]
func expandAroundCenter(s string, left,right int)(int,int){
	for ; left >= 0 && right < len(s); left,right = left-1,right+1{
		if s[left] != s[right]{
			break
		}
	}

	return left+1,right-1
}



func main(){

	//s := "asa"
	s := "sdsdashbssb"
	fmt.Println(longestPalindromeWay3(s))


}
