func repeatedCharacter(s string) byte {
    seen:=make(map[byte]bool)

    for i:=0;i<len(s);i++{
        if seen[s[i]]{
            return s[i]
        }

        seen[s[i]]=true
    }
    return 0
}