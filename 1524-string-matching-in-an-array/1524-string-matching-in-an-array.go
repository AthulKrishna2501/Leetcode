func stringMatching(words []string) []string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		for j :=0; j < len(words); j++ {
			if i!=j && strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
                break
			}
		}
	}
	return result
}