package concatention_of_array

func getConcatentation(nums []int) []int {
	return append(nums, nums...)
}
