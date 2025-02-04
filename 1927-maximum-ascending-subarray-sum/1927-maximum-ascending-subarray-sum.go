func maxAscendingSum(nums []int) int {
	n := len(nums)
	sum := nums[0]
	maxSum := 0

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			sum += nums[i]
			fmt.Println(sum)
		} else {
			if sum > maxSum {
				maxSum = sum
			}

			sum = nums[i]
		}
	}

	if sum > maxSum {
		maxSum = sum
	}

	return maxSum
}