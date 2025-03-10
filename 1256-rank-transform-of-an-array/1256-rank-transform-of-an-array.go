func arrayRankTransform(arr []int) []int {
    if len(arr)==0{
        return arr
    }

    srtArr:=append([]int{},arr...)
    sort.Ints(srtArr)
    hashMap:=make(map[int]int)
    rank:=1
    for _, num:= range srtArr{
        if _, exists:=hashMap[num];!exists{
            hashMap[num]=rank
            rank++
        }
    }

    for i,num:= range arr{
        arr[i]=hashMap[num]
    }

    return arr
}