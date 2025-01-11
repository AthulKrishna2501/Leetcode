func findNonMinOrMax(nums []int) int {
    if len(nums)<=2{
        return -1
    }

    sort.Ints(nums)
    min:=nums[0]
    max:=nums[len(nums)-1]
    for _, num:= range nums{
        if num!=min && num!=max{
            return num
        }
    }
    return -1
}