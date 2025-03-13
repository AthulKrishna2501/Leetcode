func findFinalValue(nums []int, original int) int {
	seen := make(map[int]bool)

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
		}
	}

	for seen[original] {
		original = 2 * original
	}

    return original

}