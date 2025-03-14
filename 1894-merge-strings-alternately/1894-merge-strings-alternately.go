func mergeAlternately(word1 string, word2 string) string {
    i,j:=0,0
    result:=[]rune{}
    for i<len(word1) || j<len(word2){
        if i<len(word1){
            result=append(result,rune(word1[i]))
            i++
        }
        if j<len(word2){
            result=append(result,rune(word2[j]))
            j++
        }
    }

    return string(result)
}