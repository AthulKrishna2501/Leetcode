func sumOfUnique(nums []int) int {
    sum:=0
    freq:=make(map[int]int)
    for _, num:= range nums{
        freq[num]++
    }

    for key,value:= range freq{
        if value==1{
            sum+=key
        }
    }

    return sum
}