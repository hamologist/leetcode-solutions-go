package permutations

func permute(nums []int) [][]int {
	var result [][]int
	var inner func([]int, []int)
	inner = func(current []int, remaining []int) {
		if len(remaining) == 0 {
			result = append(result, current)
		}

		for i, v := range remaining {
			inner(
				append(current, v),
				append(append([]int{}, remaining[:i]...), remaining[i+1:]...),
			)
		}
	}
	inner([]int{}, nums)

	return result
}
