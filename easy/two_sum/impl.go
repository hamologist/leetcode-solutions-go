package two_sum

func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)

	for i, num := range nums {
		if index, ok := lookup[target-num]; ok {
			return []int{index, i}
		}

		lookup[num] = i
	}

	return []int{-1, -1}
}
