package product_of_array_except_self

import (
	"slices"
)

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1
	for index, num := range nums {
		result[index] = prefix
		prefix *= num
	}

	postfix := 1
	for index, num := range slices.Backward(nums) {
		result[index] *= postfix
		postfix *= num
	}

	return result
}
