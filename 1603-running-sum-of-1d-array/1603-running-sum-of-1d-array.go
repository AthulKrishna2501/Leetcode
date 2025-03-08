func runningSum(nums []int) []int {
	result:=make([]int,len(nums))

    sum:=0
    for i:=0;i<len(nums);i++{
        sum+=nums[i]
        result[i]=sum
    }

    return result
}