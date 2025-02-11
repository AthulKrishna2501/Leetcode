func reverseWords(s string) string {
    reverseWrd:=[]string{}
    words:=strings.Split(s," ")
    for _, str:= range words{
        reverseWrd=append(reverseWrd,Reverse(string(str)))
    }

    return strings.Join(reverseWrd," ")
}

func Reverse(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}


	return string(runes)
}