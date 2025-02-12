func stableMountains(height []int, threshold int) []int {
	result := []int{}
	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			result = append(result, i)
		}
	}
    return result
}