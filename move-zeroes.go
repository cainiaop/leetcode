package main

import "fmt"

func main() {
	data := []int{0,1,0,3,12}

	moveZeroes(data)

	//Output: [1,3,12,0,0]

}

func moveZeroes(nums []int)  {

	count := len(nums)
	if count < 2 {
		return
	}

	targetIndex := 0
	for i := 0; i < count; i++ {

		if nums[i] != 0 {
			if targetIndex != i {
				nums[targetIndex] = nums[i]
				nums[i] = 0
			}

			targetIndex++
		}
	}

	fmt.Println(nums)
}