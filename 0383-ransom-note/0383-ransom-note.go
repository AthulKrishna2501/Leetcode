func canConstruct(ransomNote string, magazine string) bool {
    hashMap:=make(map[rune]int)
    for _, str:= range magazine{
        hashMap[str]++
    }

    for _, str2:= range ransomNote{
        if hashMap[str2]==0{
            return false
        }
        hashMap[str2]--
    }
    return true

}