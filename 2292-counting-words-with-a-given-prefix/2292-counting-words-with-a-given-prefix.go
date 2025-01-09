func prefixCount(words []string, pref string) int {
    count:=0
    for _, str:= range words{
        if strings.HasPrefix(str,pref){
            count++
        }
    }
    return count
}
