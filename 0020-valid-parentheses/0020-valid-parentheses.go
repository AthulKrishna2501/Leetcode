func isValid(s string) bool {
    hashMap:=map[rune]rune{
        ')':'(',
        '}':'{',
        ']':'[',
    }

    var stack []rune

    for _, char := range s{
        if char=='('|| char=='{'|| char=='['{
            stack=append(stack,char)
        }else{
            if len(stack)==0 || stack[len(stack)-1]!=hashMap[char]{
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack)==0

}