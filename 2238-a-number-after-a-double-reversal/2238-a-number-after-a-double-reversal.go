func isSameAfterReversals(num int) bool {
    if num==0{
        return true
    }

	str := strconv.Itoa(num)
	isSame := Reverse(str)

    return isSame

}

func Reverse(str string) bool {
	rvrStr := ""
	finalStr := ""

	for i := len(str) - 1; i >= 0; i-- {
		rvrStr += string(str[i])
	}

	for string(rvrStr[0]) == "0" {
		rvrStr = rvrStr[1:]
	}

	for i := len(rvrStr) - 1; i >= 0; i-- {
		finalStr += string(rvrStr[i])
	}

	if finalStr != str {
		return false
	}

	return true

}
