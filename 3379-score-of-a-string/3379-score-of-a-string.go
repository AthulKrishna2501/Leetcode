func scoreOfString(s string) int {
	count := 0.0
	for i := 0; i < len(s)-1; i++ {
		count += math.Abs(float64(s[i]) - float64(s[i+1]))
	}
    return int(count)
}