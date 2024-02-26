package replace_elements_with_greatest_element_on_right_side

import "math"

func replaceElements(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	currentGreatest := -1

	for i := len(result) - 1; i >= 0; i-- {
		nextGreatest := int(math.Max(float64(arr[i]), float64(currentGreatest)))
		result[i] = currentGreatest
		currentGreatest = nextGreatest
	}

	return result
}
