package main

import "fmt"

func main() {
	data := []int{0,1,2,3,4,10}

	//	half := map[int]int{10:0,9:1,8:2,7:3,6:4,0:5}

	rotate(data, 12)
}


func rotate(nums []int, k int)  {

	count := len(nums)
	if count < 2 || k % count == 0 {
		return
	}
	if k > count {
		k = k - count
	}

	var tmpIndex int
	tmpNums := make([]int, count)
	_ = copy(tmpNums, nums)

	for i := count - 1; i >= 0; i-- {
		tmpIndex = i - k
		if tmpIndex < 0 {
			tmpIndex = count + tmpIndex
		}

		nums[i] = tmpNums[tmpIndex]
	}

	fmt.Println(nums)
}
