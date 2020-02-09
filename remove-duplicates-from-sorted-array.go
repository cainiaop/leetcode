package main

import (
	"fmt"
)

func main() {

	data := []int{0,0,1,1,1,2,2,3,3,4}

	fmt.Println(removeDuplicates(data))
}

func removeDuplicates(nums []int) int {

	//1. 判断前面数组是否包含当前值 TODO 没考虑顺序问题 next
	//2. 如果包含，清除该值，并用于后面值向前迁移
	//3. 如果不包含，不做任何处理

	count := len(nums)
	if count <= 1 {
		return count
	}

	//tmpValue := nums[0]
	//
	//for i < len(nums) {
	//
	//	1. //if inArray(nums, i) == true {
	//	//	nums = append(nums[:i], nums[i+1:]...)
	//	//	i--
	//	//}
	//
	//	2. if nums[i] == tmpValue {
	//		nums = append(nums[:i], nums[i+1:]...)
	//		i--
	//	}
	//
	//
	//	i++
	//}

	// 3.
	// lastCompareValue := nums[0] //最后比较的值
	//targetPos := 1
	//for i := 1; i < l; i++ {
	//
	//	if nums[i] != lastCompareValue {
	//		nums[targetPos] = nums[i]
	//		targetPos++
	//	}
	//
	//	lastCompareValue = nums[i]
	//}
	//
	//return targetPos

	//4.
	//repeatedCount := 0; //重复数量
	//for i := 1; i < l; i++ {
	//
	//	if nums[i] == nums[i-1] {
	//		repeatedCount++
	//	} else {
	//		nums[i-repeatedCount] = nums[i]
	//	}
	//}
	//
	//fmt.Println(nums)
	//return repeatedCount

	//5
	nonRepeatedIndex := 0 //去重后的最小下标
	for i := 1; i < count; i++ {
		if nums[i] != nums[i-1] {
			nonRepeatedIndex++
			nums[nonRepeatedIndex] = nums[i]
		}
	}

	return nonRepeatedIndex + 1
}

//无序数组
func inArray(nums []int, pos int) bool {

	flag := false
	for i := 0; i < pos; i++ {
		if nums[i] == nums[pos] {
			flag = true
			break
		}
	}

	return flag
}

