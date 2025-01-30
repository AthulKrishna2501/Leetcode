func numberOfSpecialChars(word string) int {
    seen:=make(map[rune]bool)
    var count int

    for _, val:= range word{
        seen[val]=true
    }

    for char:= range seen{
        if _, exists:= seen[char-32];exists{
            count++
        }
    }
    return count
}