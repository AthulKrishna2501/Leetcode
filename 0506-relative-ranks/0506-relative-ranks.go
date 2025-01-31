func findRelativeRanks(score []int) []string{
   n := len(score)
	answer := make([]string, n) 
	helpArr := append([]int{}, score...)

	sort.Sort(sort.Reverse(sort.IntSlice(helpArr)))

	hashMap := make(map[int]string)
	for i, val := range helpArr {
		switch i {
		case 0:
			hashMap[val] = "Gold Medal"
		case 1:
			hashMap[val] = "Silver Medal"
		case 2:
			hashMap[val] = "Bronze Medal"
		default:
			hashMap[val] = strconv.Itoa(i + 1) 
		}
	}

	for i, val := range score {
		answer[i] = hashMap[val]
	}

	return answer
}