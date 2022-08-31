package main

import (
	"fmt"
	"sort"
)


//排序+双指针法思路：对于三层循环的优化
//nums有序保证了循环得到的符合条件的三个数间顺序从小到大，避免重复出现（-1，-1，2）和(2,-1,-1)的情况
//循环确定nums[i]后，对于nums[j]和nums[z]，可以利用双指针的思路（要保证最终相加值不变，需要一个递增，一个递减）
//leetcode链接：https://leetcode.cn/problems/3sum/
func getThreeSum(nums []int)[][]int{

	l := len(nums)
	//排序
	sort.Ints(nums)

	var dsts [][]int
	for i := 0; i < l; i++{
		//保证相等的数不需要进行重复遍历（这里有序的前提也保证了这点）
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		z := l-1
		for j := i+1; j < l; j++{
			//保证相等的数不需要进行重复遍历（这里有序的前提也保证了这点）
			if j > i+1 && nums[j] == nums[j-1]{
				continue
			}
			//这里z不需要每次从l-1位置开始，因为nums[j]增大了，nums[z]必定要比上一次更小，位置可以从上次循环结束位置开始
			//而且当nums[j]+nums[z]<-1*nums[i]时，不进行遍历的原因在于继续遍历下去也不会有符合条件的z，因为nums[z]是递减的，nums[j]+nums[z]只会减小
			for ; z > j && nums[j]+nums[z] >= -1*nums[i]; z--{
				if nums[i]+nums[j]+nums[z] == 0{
					dst := []int{nums[i],nums[j],nums[z]}
					dsts = append(dsts,dst)
					break
				}
			}
		}
	}

	return dsts
}

func main() {

	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println(getThreeSum(nums))

}
