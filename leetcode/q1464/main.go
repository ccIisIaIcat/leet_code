package main

func maxProduct(nums []int) int {
	max_list := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] - 1
	}
	max_list[0] = nums[0]
	max_list[1] = nums[1]
	for i := 2; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		temp_position, temp_value := get_min(max_list)
		if nums[i] > temp_value {
			max_list[temp_position] = nums[i]
		}
	}
	return max_list[0] * max_list[1]

}

func get_min(list_ []int) (int, int) {
	if list_[0] <= list_[1] {
		return 0, list_[0]
	} else {
		return 1, list_[1]
	}
}
