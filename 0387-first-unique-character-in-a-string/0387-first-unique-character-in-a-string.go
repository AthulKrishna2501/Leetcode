func firstUniqChar(s string) int {
	freq := make(map[string]int)

	for _, str := range s {
		freq[string(str)]++
	}


    for i,char:= range s{
        if freq[string(char)]==1{
            return i
        }
    }

	// firstStr := ""
	// for key, value := range freq {
	// 	if value == 1 {
	// 		firstStr = key
	// 		break
	// 	}
	// }

	// for i, char := range s {
	// 	if string(char) == firstStr {
	// 		return i
	// 	}
	// }

	return -1
}
