package majority_element

func majorityElement(nums []int) int {
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num] += 1

		if counter[num] > len(nums)/2 {
			return num
		}
	}

	return -1
}
