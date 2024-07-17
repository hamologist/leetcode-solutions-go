package contains_duplicate

func containsDuplicateRedo(nums []int) bool {
	lookup := make(map[int]struct{})

	for _, num := range nums {
		_, ok := lookup[num]
		if ok {
			return true
		}

		lookup[num] = struct{}{}
	}

	return false
}
