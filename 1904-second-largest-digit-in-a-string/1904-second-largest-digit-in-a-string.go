func secondHighest(s string) int {
    arr:=[]int{}
    seen:=make(map[int]bool)

    for _, str:= range s{
        if !unicode.IsLetter(str){
            num,_:=strconv.Atoi(string(str))
            if !seen[num]{
            arr=append(arr,num)
            seen[num]=true
            }
        }
    }

    if len(arr)<=1{
        return -1
    }

    sort.Ints(arr)

    return arr[len(arr)-2]
}