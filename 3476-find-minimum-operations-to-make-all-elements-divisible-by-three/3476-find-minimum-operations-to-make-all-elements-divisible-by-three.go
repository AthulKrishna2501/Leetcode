func minimumOperations(nums []int) int {
    count:=0
    for _, num:= range nums{
        if (num+1)%3==0 || (num-1)%3==0{
            count++
        }
    }
    return count
}