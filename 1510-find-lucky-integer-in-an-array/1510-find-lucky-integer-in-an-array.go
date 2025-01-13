func findLucky(arr []int) int {
    hashMap:=make(map[int]int)

    for _, num:= range arr{
        hashMap[num]++
    }

    result:=[]int{}
    for key,value:= range hashMap{
        if key==value{
            result=append(result,key)
        }
    }

    if len(result)==1{
        return result[0]
    }else if len(result)==0{
        return -1
    }

    maxNum:=0
    for _, num:= range result{
        if num>maxNum{
            maxNum=num
        }
    }

    return maxNum
}