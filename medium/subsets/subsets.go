package subsets

func subsets(nums []int) [][]int {
	var result [][]int

	var inner func([]int, []int)
	inner = func(current, remaining []int) {
		result = append(result, current)

		for i, num := range remaining {
			inner(
				append(append([]int{}, current...), num),
				append([]int{}, remaining[i+1:]...),
			)
		}
	}
	inner([]int{}, nums)

	return result
}
