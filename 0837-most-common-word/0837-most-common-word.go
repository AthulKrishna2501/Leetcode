func mostCommonWord(paragraph string, banned []string) string {
	hashMap := make(map[string]int)
	isBanned := make(map[string]bool)
    maxWord:=""
    count:=0

	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return !unicode.IsLetter(r)
	})


	for _, char := range banned {
		isBanned[char] = true
	}

	for _, str := range words {
		lower := strings.ToLower(str)
		if !isBanned[lower] {
			hashMap[lower]++
		}
	}

    for key,value:=range hashMap{
        if value>count{
            count=value
            maxWord=key
        }
    }

    return maxWord

	
}