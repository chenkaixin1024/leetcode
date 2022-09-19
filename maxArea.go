package main

import "fmt"

//题目：盛水最多的容器
//leetcode链接：https://leetcode.cn/problems/container-with-most-water/

//双指针法思路：
//数组内的元素作为纵轴高度，与横轴构成的面积=左右两个纵轴高度的较小值*横轴距离
//因此寻找最大面积，需要在横轴距离减小的情况下，改变左右纵轴高度的较小值才有可能获得更大的面积
//到此，可以利用左右两侧指针分别指向数组两端，将较小值向对侧偏移，去寻找较大的面积，直到两个指针相遇
func findMaxArea(height []int) int{

	l := len(height)
	left := 0
	right := l-1
	max := 0

	for left < right{
		area := 0
		if height[left] <= height[right]{
			area = height[left] * (right - left)
			left++
		}else{
			area = height[right] * (right - left)
			right--
		}

		if area > max{
			max = area
		}
	}

	return max
}

func main(){

	//height := []int{1,8,6,2,5,4,8,3,7}
	height := []int{1,1}
	fmt.Println(findMaxArea(height))
}
