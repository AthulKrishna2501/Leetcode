func isAnagram(s string, t string) bool {
    seen:=make(map[rune]int)

    if len(s)!=len(t){
        return false
    }

    for _,str:= range s{
        seen[str]++
    }

    for _, str2:= range t{
        if _, found:=seen[str2];!found{
            return false
        }

        seen[str2]--
        if seen[str2]==0{
            delete(seen,str2)
        }
    }

    return true

}