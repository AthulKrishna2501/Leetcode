func uncommonFromSentences(s1 string, s2 string) []string {
    result:=[]string{}
    freq:=make(map[string]int)

    split1:=strings.Split(s1," ")
    split2:=strings.Split(s2," ")

    for _,str1:= range split1{
        freq[string(str1)]++
    }

    for _, str2:= range split2{
        freq[string(str2)]++
    }

    for key,value:= range freq{
        if value==1{
            result=append(result,key)
        }
    }

    return result
}