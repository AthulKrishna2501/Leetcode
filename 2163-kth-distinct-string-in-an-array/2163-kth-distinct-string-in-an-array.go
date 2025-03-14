func kthDistinct(arr []string, k int) string {
    freq:=make(map[string]int)

    isKth:=0
    for _,str:= range arr{
        freq[string(str)]++
    }


    for _,str:= range arr{
        if freq[str]==1{
            isKth++
            if isKth==k{
                return str
            }
        }
    }

    return ""
}