package two_sum

func twoSumRedo(nums []int, target int) []int {
	lookup := make(map[int]int)

	for index, num := range nums {
		lookupIndex, ok := lookup[target-num]

		if ok {
			return []int{lookupIndex, index}
		}

		lookup[num] = index
	}

	return []int{-1, -1}
}
