func countPrefixes(words []string, s string) int {
    count:=0
    for _, str:=range words{
        if strings.HasPrefix(s,str){
            count++
        }
    }

    return count
}