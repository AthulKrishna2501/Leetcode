func maximumNumberOfStringPairs(words []string) int {

	rvrArr := []string{}

	count := 0
	for _, str := range words {
		rvrStr := reverse(str)
		rvrArr = append(rvrArr, rvrStr)

	}

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(rvrArr); j++ {
			if words[i] == rvrArr[j] {
				count++
			}
		}
	}

    return count
}

func reverse(str string) string {
	runes := []rune(str)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}