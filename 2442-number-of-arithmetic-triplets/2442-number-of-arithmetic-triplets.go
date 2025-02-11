func arithmeticTriplets(nums []int, diff int) int {
    count:=0
    for i:=0;i<len(nums)-1;i++{
        for j:=i+1;j<len(nums)-1;j++{
            for k:=j+1;k<len(nums);k++{
                if i!=k && i!=j && i<j && j<k && nums[j]-nums[i]==diff && nums[k]-nums[j]==diff{
                    count++
                }
            }
        }
    }
    return count
}