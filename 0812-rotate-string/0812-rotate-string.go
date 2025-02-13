

func rotateString(s string, goal string) bool {
	split := strings.Split(s, "")

	for i := 0; i < len(s); i++ {

		check := split[0]
		split = split[1:]
		split = append(split, check)
		str := strings.Join(split, "")
		if str == goal {
            return true

		}
	}

    return false

}