func areOccurrencesEqual(s string) bool {
	freq := make(map[string]int)

	freqArr := []int{}

	for _, str := range s {
		freq[string(str)]++
	}

	for _, value := range freq {
		freqArr = append(freqArr, value)
	}

	for i := 0; i < len(freqArr)-1; i++ {
		if freqArr[i] != freqArr[i+1] {
			return false
		}
	}

	return true

}