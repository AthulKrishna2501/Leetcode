func findRestaurant(list1 []string, list2 []string) []string {
	hashMap := make(map[string]int)
	result := []string{}
	minIndexSum := math.MaxInt

	for i, str := range list1 {
		hashMap[str] = i
	}

	for j, str := range list2 {
		if i, exists := hashMap[str]; exists {
			indexSum := i + j
			if indexSum < minIndexSum {
				minIndexSum = indexSum
				result = []string{str}
			} else if indexSum == minIndexSum {
				result = append(result, str)
			}
		}
	}

	return result

}