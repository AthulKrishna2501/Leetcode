func findMissingAndRepeatedValues(grid [][]int) []int {
    hashMap := make(map[int]int)
	arr := []int{}

	for _, nums := range grid {
		for _, num := range nums {
			hashMap[num]++
			arr = append(arr, num)
		}
	}

	sort.Ints(arr)

	repeated := -1
	missing := -1
	n := len(arr) 

	for i := 1; i <= n; i++ {
		if hashMap[i] == 2 {
			repeated = i
		} else if hashMap[i] == 0 {
			missing = i
		}
	}

    return []int{repeated,missing}
}