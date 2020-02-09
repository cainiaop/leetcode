package main

import "fmt"

func main() {
	data := []int{0,1,2,3,4,10}

//	half := map[int]int{10:0,9:1,8:2,7:3,6:4,0:5}

	fmt.Println(twoSum(data, 10))

}


func twoSum(nums []int, target int) []int {

	count := len(nums)
	if count < 2 {
		return nil
	}

	half := make(map[int]int, count)
	var result []int

	for i := 0; i < count; i++ {

		v, ok := half[nums[i]]
		if ok {
			result = append(result, v, i)
			break
		} else {
			half[target - nums[i]] = i
		}
	}

	return result
}
