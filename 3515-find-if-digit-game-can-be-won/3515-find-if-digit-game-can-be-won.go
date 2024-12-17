func canAliceWin(nums []int) bool {
    sum:=0
    count:=0
    for _, num := range nums{
        if num<=9{
            count+=num
        }
        if num>=10{
            sum+=num
        }
    }
    if sum==count{
        return false
    }
    return true
}