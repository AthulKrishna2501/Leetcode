func longestCommonPrefix(strs []string) string {
    temp:=findMin(strs)
    for i,s:= range temp{
        for _,str:= range strs{
            if byte(s)!=str[i]{
                return temp[:i]
            }
        }
    }
    return temp
}

func findMin(str []string)string{
    min:=str[0]
    for _, value:= range str{
        if len(value)<len(min){
            min=value
        }
    }

    return min
}