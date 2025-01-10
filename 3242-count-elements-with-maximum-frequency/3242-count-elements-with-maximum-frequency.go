func maxFrequencyElements(nums []int) int {
    hashMap:=make(map[int]int)
    for _, num:= range nums{
        hashMap[num]++
    }

    maxCount:=0
    count:=0
    for _,value :=range hashMap{
        if value>maxCount{
            maxCount=value
        }
    }

    for _, value:= range hashMap{
        if value==maxCount{
            count+=value
        }
    }

    return count



}