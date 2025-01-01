package top_k_frequent_elements

import "slices"

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, num := range nums {
		counter[num] += 1
	}

	for num, count := range counter {
		freq[count] = append(freq[count], num)
	}

	result := make([]int, k)
	for _, nums := range slices.Backward(freq) {
		for _, num := range nums {
			result[k-1] = num

			k -= 1
			if k <= 0 {
				return result
			}
		}
	}

	return result
}
