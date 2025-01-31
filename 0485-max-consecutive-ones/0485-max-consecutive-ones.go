func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			if count > maxCount {
				maxCount = count
			}
            count = 0
		}
	}

	if count > maxCount {
        maxCount=count
	}

    return maxCount
}