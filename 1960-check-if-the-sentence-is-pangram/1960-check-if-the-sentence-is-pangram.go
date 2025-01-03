func checkIfPangram(sentence string) bool {
    seen:=make(map[rune]bool)
    for _, str:= range sentence{
        if !seen[str]{
            seen[str]=true
        }
    }

    return len(seen)==26
}