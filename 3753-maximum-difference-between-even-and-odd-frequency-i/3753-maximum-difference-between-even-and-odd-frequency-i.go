func maxDifference(s string) int {
	count := make(map[string]int)

	for _, str := range s {
		count[string(str)]++
	}

	odd := 0
	even := math.MaxInt
	for _, value := range count {
		if value%2 == 0 {
			if value <even {
				even = value
			}
		} else if value > odd {
			odd = value
		}
	}

    return odd-even
}