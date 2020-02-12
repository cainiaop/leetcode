package main

import (
	"fmt"
)

func main() {
	//data := []int{1,2,3,4,10}

	data := []int{1,2,3,4,5,6}
	rotate(data, 2)
}


func rotate(nums []int, k int)  {

	count := len(nums)
	if count < 2 || k < 1 {
		return
	}
	k %= count
	if k == 0 {
		return
	}

	// 1：申请单独内存存储原有nums
	//var tmpIndex int
	//tmpNums := make([]int, count)
	//_ = copy(tmpNums, nums)
	//
	//for i := count - 1; i >= 0; i-- {
	//	tmpIndex = i - k
	//	if tmpIndex < 0 {
	//		tmpIndex = count + tmpIndex
	//	}
	//
	//	nums[i] = tmpNums[tmpIndex]
	//}

	//2. 申请单独内存存储
	//tmpCopy := make([]int, count)
	//_ = copy(tmpCopy, nums)
	//
	//for i := 0; i < count; i++ {
	//
	//	nums[(i + k) % count] = tmpCopy[i]
	//}

	//3. 环形 暂且没搞定 @TODO
	// count % k == 0
	// start 0 -> k -> 2k ->... 0 ;
	// 1 -> 1 + k -> 1 + 2k ->... 1;
	// ....
	// k-1 -> 2k -1 -> 3k - 1 ->... k-1;
	// k stop
	//
	// count % k != 0
	// start 0 -> k -> 2k ->... 1 -> 1 + k -> 1 + 2k ->... 2 -> 2 + k -> 2 + 2k ->... k stop

	//var tmpValue int
	//start, prevIndex := 0, 0
	//prev := nums[prevIndex]
	//
	//isMultiCycle := false
	//if k != 1 && count % k == 0 {
	//	isMultiCycle = true
	//}
	//
	//fmt.Println(isMultiCycle)
	//
	//for {
	//	//fmt.Println(prevIndex, start)
	//
	//	nextIndex := (prevIndex + k) % count
	//	tmpValue = nums[nextIndex]
	//	nums[nextIndex] = prev
	//
	//	prev = tmpValue
	//	prevIndex = nextIndex
	//
	//	if isMultiCycle == true {  //多次循环
	//		fmt.Println(prevIndex, start, prev)
	//		if prevIndex == start {
	//			prevIndex += 1
	//			start += 1
	//
	//			if prevIndex == k {
	//				break
	//			}
	//		}
	//
	//	} else {	//一次循环
	//
	//		if prevIndex == start {
	//			break
	//		}
	//	}
	//}


	//4. 反转3次
	// 这个方法基于这个事实：当我们旋转数组 k 次， k\%nk%n 个尾部元素会被移动到头部，剩下的元素会被向后移动。
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])

	fmt.Println(nums)
}

func reverse(nums []int) {

	for i, j := 0, len(nums) - 1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
