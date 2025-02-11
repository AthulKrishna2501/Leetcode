func removeTrailingZeros(num string) string {
	// num = strings.TrimRight(num,"0")
	// return num

	for string(num[len(num)-1]) == "0" {
		num = num[:len(num)-1]
	}

    return num
}