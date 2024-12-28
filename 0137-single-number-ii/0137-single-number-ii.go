func singleNumber(nums []int) int {
    count:=make(map[int]int)
    for _, num:= range nums{
        count[num]++
    }

    for key,value:= range count{
        if value==1{
            return key
        }
    }
    return -1
}