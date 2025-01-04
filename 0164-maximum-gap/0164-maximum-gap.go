func maximumGap(nums []int) int {
    sort.Ints(nums)
    if len(nums)<2{
        return 0
    }

    maxDiff:=0
    for i:=0;i<len(nums)-1;i++{
        if nums[i+1]-nums[i]>maxDiff{
            maxDiff=nums[i+1]-nums[i]
        }
    }
return maxDiff

}